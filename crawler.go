package diamonds

import (
	"fmt"
	"io"
)

const DefaultBaseUrl = "http://diamondse.info/webService.php"

type Crawler struct {
	// BaseURL where the diamonds search engine can be found
	BaseURL string
	// SearchParameters are the Parameters to search the diamonds
	// database with
	SearchParameters Parameters
	// OutputStream is a Writer that takes the output of
	// of the Crawler
	OutputStream io.Writer
	// rowNumberGenerator takes a starting row, ending row, and
	// number of rows and returns a slice numbers of rows to start
	// getting pages
	rowNumberGenerator func(start, end, numRows int) []int
}

// NewCrawler returns an instance of a Crawler.
func NewCrawler(outputStream io.Writer) Crawler {
	return Crawler{
		SearchParameters:   NewParameters(),
		BaseURL:            DefaultBaseUrl,
		OutputStream:       outputStream,
		rowNumberGenerator: DefaultRowNumberGenerator,
	}
}

// Sets the rowNumberGenerator function. This function is used to generate
// row numbers used by the crawler
func (c *Crawler) SetRowNumberGenerator(rng func(start, end, numRows int) []int) {
	c.rowNumberGenerator = rng
}

// Crawl starts the crawler and prints the results to the OutPutStream
func (c Crawler) Crawl(rowNum int) error {
	done := make(chan struct{})
	defer close(done)

	start := 0
	end := 100
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

// DefaultRowNumberGenerator returns a slice of integers
// between start and end of length rowNums
func DefaultRowNumberGenerator(start, end, rowNums int) []int {
	if rowNums > (end - start) {
		panic("number of rows cannot be greater than number of rows between start and end")
	}

	out := []int{}
	for ; start < end && rowNums >= 0; start, rowNums = start+1, rowNums-1 {
		out = append(out, start)
	}

	return out
}
