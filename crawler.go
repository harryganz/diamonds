package diamonds

import (
	"fmt"
	"io"
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
	pageGetter func(baseUrl string, params Parameters, rowStart int) (io.Reader, error)
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
	}
}

// SetRowNumberGenerator sets the rowNumberGenerator function. This function is used to generate
// row numbers used by the crawler
func (c *Crawler) SetRowNumberGenerator(rng func(start, end, numRows int) []int) {
	c.rowNumberGenerator = rng
}

// SetPageGetter sets the pageGetter function. This function is used to
// get the diamond search engine results pages
func (c *Crawler) SetPageGetter(pg func(string, Parameters, int) (io.Reader, error)) {
	c.pageGetter = pg
}

// Crawl starts the crawler and prints the results to the OutPutStream
func (c Crawler) Crawl() error {
	done := make(chan struct{})
	defer close(done)

	start := 0
	end := 100
	rowNum := c.NumResults
	gen := func(done <-chan struct{}) <-chan int {
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

	for v := range gen(done) {
		fmt.Fprint(c.OutputStream, v)
	}

	return nil
}
