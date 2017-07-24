package diamonds

import (
	"io"

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
	// rowNumberGenerator takes a starting row, ending row, and
	// number of rows and returns a slice numbers of rows to start
	// getting pages
	rowNumberGenerator func(start, end, numRows int) []int
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
		rowNumberGenerator: DefaultRowNumberGenerator,
		pageGetter:         DefaultPageGetter,
		pageParser:         DefaultPageParser,
	}
}

// SetRowNumberGenerator sets the rowNumberGenerator function. This function is used to generate
// row numbers used by the crawler
func (c *Crawler) SetRowNumberGenerator(rng func(start, end, numRows int) []int) {
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

	start := 0
	end := 100
	rowNum := c.NumResults
	genNums := func(done <-chan struct{}) <-chan int {
		out := make(chan int)

		go func() {
			defer close(out)
			for _, v := range c.rowNumberGenerator(start, end, rowNum) {
				select {
				case out <- v:
				case <-done:
					return
				}
			}
		}()

		return out
	}
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

	nums := genNums(done)
	pages := genPages(done, nums)
	diamonds := parsePages(done, pages)

	var i int
	for v := range diamonds {
		if i == 0 {
			gocsv.Marshal(v, c.OutputStream)
		} else {
			gocsv.MarshalWithoutHeaders(v, c.OutputStream)
		}

		i++
	}

	return nil
}
