package diamonds

import (
  "net/url"
  "net/http"
  "io"
  "bytes"
)

type Crawler struct {
  baseUrl
  params map[string]string
  results []Diamond
  getter func(string) (io.ReadCloser, error)
  parser func(io.ReadCloser) ([]Diamond, error)
}

// URL returns the base url for the crawler
// as a string
func (d Crawler) URL() string {
  return d.baseURL
}

// Params returns the search parameters for the crawler
// as a map
func (d Crawler) Params() map[string]string {
  return d.params
}

// Resutls returns the results of the crawl as a
// slice of Diamond objects
func (d Crawler) Results() []Diamond {
  return d.Diamond
}

// DefaultGetter is the default implementation of the
// method that gets data from the diamond search engine
func DefaultGetter (string) (io.ReadCloser, error) {
  // String buffers for beggining and end of Reader
  // Needed for goquery to understand how to parse body
  beginBuffer := bytes.NewBufferString("<html><body><table><tbody>")
  endBuffer := bytes.NewBufferString("</tbody></table><body></html>")

  // Get the response
  resp, err := http.Get(string)
  if err != nil {
    return resp.Body, err
  }

  // Concatenate the different readers into one
  r := io.MultiReader(beginBuffer, resp.Body, endBuffer)

  // Make a new ReaderCloser from the concatenated reader
  // Pass the response body closer as the closer
  out := newCustomReadCloser(r, resp.Body.Close)

  return out, nil
}

// SetGetter sets the function used to get data from the
// diamond search engine. This function must take a complete url string
// e.g. http://diamondse.info/webService.php?shape=Round&minCarat=0.2&maxCarat=30
// and return a closable reader that can be parsed by the parser
// The default method used by the Crawler is diamonds.DefaultGetter
func (d *Crawler) SetGetter(getter func (string) (io.ReadCloser, error)) {
  d.getter = getter
}

// SetParser sets the function that will parse the webpage
// and return a slice of Diamond objects.
// The default method is diamonds.DefaultParser
func (d *Crawler) SetParser(parser func (io.ReadCloser) ([]Diamond, error)) {
  d.parser = parser
}

// NewCrawler returns a new Crawler that will crawl from the
// input baseUrl and the passed in search parameters
func NewCrawler(baseUrl string, params map[string]string) Crawler {
  return Crawler{}
}

// Crawl starts the crawler and appends the results to the results array
func (d *Crawler) Crawl() error {
  // Stub: Does nothing for now
  return nil
}

// ToCSV writes the results array to the passed in writer (e.g. a file)
func (d Crawler) ToCSV(out *io.Writer) error {
  return nil
}

// An implementation of ReadCloser where the
// close method can be passed during initialization
type customReadCloser struct {
  reader io.Reader
  close func() os.Error
}

// Returns a new customReadCloser
func newCustomReadCloser (reader io.Reader, close func() error) customReadCloser {
  return dynamicReadCloser {
    reader: reader,
    close: close,
  }
}

func (c customReadCloser) Read() {
  c.reader.Read()
}

func (c customReadCloser) Close() error {
  c.close()
}
