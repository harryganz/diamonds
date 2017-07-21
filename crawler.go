package diamonds

import (
  "io"
)

const DefaultBaseUrl = "http://diamondse.info/webService.php"

type Crawler struct {
  // The BaseUrl where the diamonds search engine can be found
  BaseUrl string
  // The starting search parameters to use
  StartParameters Parameters
  pageGenerator func(done <-chan struct{}, start, end, n int) (<-chan int)
  pageGetter func(done <-chan struct{}, params Parameters, pageNum <-chan int) (<-chan io.ReadCloser, <-chan error)
  pageParser func(done <-chan struct{}, page <-chan io.ReadCloser) (<-chan []Diamond, <-chan error)
}

func NewCrawler() Crawler {
  return Crawler {
    BaseUrl: DefaultBaseUrl,
    StartParameters: NewParameters(),
    pageGenerator: pageGenerator,
    pageGetter: pageGetter,
    pageParser: pageParser,
  }
}
