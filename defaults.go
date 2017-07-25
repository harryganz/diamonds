package diamonds

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
  "time"
  "math/rand"
  "regexp"

	"github.com/PuerkitoBio/goquery"
)

// DefaultBaseURL is the url for the diamonds search engine database
const DefaultBaseURL = "http://diamondse.info/webService.php"

// DefaultRowNumberGenerator returns a slice of integers
// between start and end of length rowNums
func DefaultRowNumberGenerator(start, end, rowNums int) []int {
	if rowNums > (end - start) {
		panic("number of rows cannot be greater than number of rows between start and end")
	}

	pages := sample(rowNums/20, start/20, end/20)
  rows := []int{}
  for _, page := range pages {
    rows = append(rows, page * 20)
  }

  return rows
}

// DefaultPageGetter returns a page from the diamond search engine using the defaults
func DefaultPageGetter(baseURL string, params Parameters, rowStart int) (io.ReadCloser, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nopReaderCloser(), err
	}

	params.SetRow(rowStart)

	q := u.Query()
	for k, v := range params.ToMap() {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nopReaderCloser(), err
	}

	head := bytes.NewBufferString("<html><body><table><tbody>")
	tail := bytes.NewBufferString("</tbody></table></body></html>")

	mr := io.MultiReader(head, resp.Body, tail)
	mrc := multiReadCloser{
		reader: mr,
		close:  resp.Body.Close,
	}

	return mrc, nil
}

// DefaultPageParser returns a slice of Diamonds
func DefaultPageParser(page io.ReadCloser) ([]Diamond, error) {
	results := []Diamond{}
	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return []Diamond{}, err
	}

	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		diamond := Diamond{}
		s.Find("td").Each(func(k int, ss *goquery.Selection) {
			text := ss.Text()

			switch k {
			case 1:
				diamond.Shape = text
			case 2:
				diamond.Carat = parseFloat(text)
			case 3:
				diamond.Cut = text
			case 4:
				diamond.Color = text
			case 5:
				diamond.Clarity = text
			case 6:
				diamond.Width = parseFloat(text)
			case 7:
				diamond.Depth = parseFloat(text)
			case 8:
				diamond.Certification = text
			case 9:
				diamond.Dimensions = text
			case 10:
				diamond.Price = parseFloat(text[1:])
			}
		})

		results = append(results, diamond)
	})

	if r := recover(); r != nil {
		return []Diamond{}, errors.New("Unknown error occurred")
	}

	return results, nil
}

// Helpers for DefaultRowNumGenerator

func sample(k, start, end int) []int {
	var r []int
	src := rand.NewSource(time.Now().Unix())
	rnd := rand.New(src)

	for i, v, n := 0, start, (end - start + 1); i < n; i, v = i + 1, v + 1 {
		if i < k {
			r = append(r, v)
		} else {
			j := rnd.Intn(n)

			if j < k {
				r[j] = v
			}
		}
	}

	return r
}

// Helpers for DefaultPageGetter

func nopReaderCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBuffer([]byte{}))
}

type multiReadCloser struct {
	reader io.Reader
	close  func() error
}

func (mrc multiReadCloser) Read(p []byte) (int, error) {
	return mrc.reader.Read(p)
}

func (mrc multiReadCloser) Close() error {
	return mrc.close()
}

// Helpers for DefaultPageParser
var commaRegex = regexp.MustCompile(",")
func parseFloat(s string) float64 {
	result, err := strconv.ParseFloat(commaRegex.ReplaceAllString(s, ""), 64)
	if err != nil {
		panic(err)
	}

	return result
}
