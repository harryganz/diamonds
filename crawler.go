package diamonds

import (
  "io"
)

const DefaultBaseUrl = "http://diamondse.info/webService.php"

type Crawler struct {
  // The BaseUrl where the diamonds search engine can be found
  baseUrl string
  // The starting search parameters to use
  startParameters Parameters
}

func NewCrawler () Crawler {
  return Crawler {
    startParameters: NewParameters(),
    baseUrl: DefaultBaseUrl,
  }
}
