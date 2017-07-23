package diamonds

import (
  "io"
)

const DefaultBaseUrl = "http://diamondse.info/webService.php"

type Crawler struct {
  // BaseUrl where the diamonds search engine can be found
  BaseUrl string
  // SearchParameters are the Parameters to search the diamonds
  // database with
  SearchParameters Parameters
  // rowNumberGenerator takes a starting row, ending row, and
  // number of rows and returns a slice numbers of rows to start
  // getting pages
  rowNumberGenerator func (start, end, numRows int) []int
  // pageGetter is a function that takes a baseUrl,
  // parameters, and rowNumber and returns an io.ReadCloser and
  // error
  pageGetter func (baseUrl string, params Parameters, rowNum int) (io.ReadCloser, error)

}

// NewCrawler returns an instance of a Crawler.
func NewCrawler () Crawler {
  return Crawler {
    StartParameters: NewParameters(),
    BaseUrl: DefaultBaseUrl,


  }
}
