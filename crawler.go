package diamonds

import (
	"bytes"
	"io"
	"io/ioutil"
	"regexp"
	"strconv"
	"sync"

	"github.com/gocarina/gocsv"
)

// Crawler is the managing struct for the diamonds scraper
// it contains a method, Crawl, that crawls the database and returns
// the results to a writer outputstream
type Crawler struct {
	// BaseURL where the diamonds search engine can be found
	BaseURL string
	// SearchParameters are the Parameters to search the diamonds
	// database with
	SearchParameters Parameters
	// OutputStream is a Writer that takes the output of
	// of the Crawler
	OutputStream io.Writer
	// The number of results to Crawl
	NumResults int
	// The number of threads to use in querying search engine
	NumThreads int
	// rowNumberGenerator takes an ending row, and
	// number of rows and returns a slice numbers of rows to start
	// getting pages
	rowNumberGenerator func(end, numRows int) []int
	// Page getter takes a baseUrl, parameters, and rowNum
	// and returns a reader and error
	pageGetter func(baseUrl string, params Parameters, rowStart int) (io.ReadCloser, error)
	// Pageparser parses the page and returns a slice of Diamonds
	pageParser func(io.ReadCloser) ([]Diamond, error)
}

// NewCrawler returns an instance of a Crawler.
func NewCrawler(numResults int, outputStream io.Writer) Crawler {
	return Crawler{
		SearchParameters:   NewParameters(),
		BaseURL:            DefaultBaseURL,
		OutputStream:       outputStream,
		NumResults:         numResults,
		NumThreads:         1,
		rowNumberGenerator: DefaultRowNumberGenerator,
		pageGetter:         DefaultPageGetter,
		pageParser:         DefaultPageParser,
	}
}

// SetRowNumberGenerator sets the rowNumberGenerator function. This function is used to generate
// row numbers used by the crawler
func (c *Crawler) SetRowNumberGenerator(rng func(end, numRows int) []int) {
	c.rowNumberGenerator = rng
}

// SetPageGetter sets the pageGetter function. This function is used to
// get the diamond search engine results pages
func (c *Crawler) SetPageGetter(pg func(string, Parameters, int) (io.ReadCloser, error)) {
	c.pageGetter = pg
}

// Crawl starts the crawler and prints the results to the OutPutStream
func (c Crawler) Crawl() error {
	done := make(chan struct{})
	defer close(done)

	firstPage, err := c.pageGetter(c.BaseURL, c.SearchParameters, 0)
	if err != nil {
		return err
	}
	// Total number of rows
	totalRows, err := getTotalRows(firstPage)
	if err != nil {
		return err
	}

	// Number of results to return (within 20)
	rowNum := c.NumResults

	// Returns row numbers to use
	genNums := func(done <-chan struct{}) <-chan int {
		out := make(chan int)

		go func() {
			defer close(out)
			for _, v := range c.rowNumberGenerator(totalRows, rowNum) {
				select {
				case out <- v:
				case <-done:
					return
				}
			}
		}()

		return out
	}

	// Returns pages
	genPages := func(done <-chan struct{}, nums <-chan int) <-chan io.ReadCloser {
		out := make(chan io.ReadCloser)

		go func() {
			defer close(out)
			for v := range nums {
				page, _ := c.pageGetter(c.BaseURL, c.SearchParameters, v)
				select {
				case out <- page:
				case <-done:
					return
				}
			}
		}()

		return out
	}

	// Merges genPages channels together for efficiency
	mergePages := func(done <-chan struct{}, cs ...<-chan io.ReadCloser) <-chan io.ReadCloser {
		out := make(chan io.ReadCloser)
		var wg sync.WaitGroup // Wait group to wait for all channels to close

		// Adds output from single channel to combined channel
		output := func(c <-chan io.ReadCloser) {
			defer wg.Done()
			for p := range c {
				select {
				case out <- p:
				case <-done:
					return
				}
			}
		}

		// Creating a waiting group for every channel
		wg.Add(len(cs))
		// Merge channels to output
		for _, c := range cs {
			go output(c)
		}

		// Wait for all channels to be done to close
		go func() {
			wg.Wait()
			close(out)
		}()

		return out
	}

	// Parse pages to get Diamond data
	parsePages := func(done <-chan struct{}, pages <-chan io.ReadCloser) <-chan []Diamond {
		out := make(chan []Diamond)

		go func() {
			defer close(out)
			for page := range pages {
				diamonds, _ := c.pageParser(page)
				page.Close()
				select {
				case out <- diamonds:
				case <-done:
					return
				}
			}
		}()

		return out
	}

	// Start procedure

	nums := genNums(done)
	// Utilize multiple threads to get pages as this is the rate limiting
	// step, Fan-Out
	var pageWorkers []<-chan io.ReadCloser
	for i := 0; i < c.NumThreads; i++ {
		pageWorkers = append(pageWorkers, genPages(done, nums))
	}
	// Merge workers into one channel Fan-In
	pages := mergePages(done, pageWorkers...)
	diamonds := parsePages(done, pages)

	// TODO use a configured text marshaller instead of hardcoding it here
	var i int
	for v := range diamonds {
		if i == 0 {
			gocsv.Marshal(v, c.OutputStream)
		} else {
			gocsv.MarshalWithoutHeaders(v, c.OutputStream)
		}

		i++
	}

	// TODO error handling (e.g. error channel with consuming go func)
	return nil
}

var rowNumRegex = regexp.MustCompile("Displaying [0-9]+ to [0-9]+ of [0-9]+ diamonds")

func getTotalRows(page io.ReadCloser) (int, error) {
	body, err := ioutil.ReadAll(page)
	if err != nil {
		return 0, err
	}
	slice := rowNumRegex.Find(body)
	str := bytes.Split(slice, []byte(" "))[5]

	return strconv.Atoi(string(str))
}
