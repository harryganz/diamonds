package diamonds

import (
	"bytes"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// DefaultBaseURL is the url for the diamonds search engine database
const DefaultBaseURL = "http://diamondse.info/webService.php"
const pageSize = 20

// DefaultRowNumberGenerator returns a slice of integers
// between start and end of length rowNums
func DefaultRowNumberGenerator(end, rowNums int) []int {
	if rowNums > end {
		panic("number of rows cannot be greater than number of total rows")
	}

	pages := sample(rowNums/20, end/20)
	rows := []int{}
	for _, page := range pages {
		rows = append(rows, page*20)
	}

	return rows
}

// DefaultPageGetter returns a page from the diamond search engine as a ReadCloser
// given a baseUrl string, a Parameter instance, and a starting row  as an int.
// It will surround the returned page with html, body, table, and tbody tags so
// the parser can parse it.
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

// DefaultPageParser parses an html page from the diamonds search engine
// and returns an slice of Diamond objects
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

	return results, nil
}

// Helpers for DefaultRowNumGenerator

// sample returns a slice random uniform sample of integers
// of length k from 0 through end as a slice of integers
func sample(k, end int) []int {
	var r []int
	src := rand.NewSource(time.Now().Unix())
	rnd := rand.New(src)

	for i, n := 0, (end + 1); i < n; i++ {
		if i < k {
			r = append(r, i)
		} else {
			j := rnd.Intn(n)

			if j < k {
				r[j] = i
			}
		}
	}

	return r
}

// Helpers for DefaultPageGetter

// Returns a no-op ReadCloser
func nopReaderCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBuffer([]byte{}))
}

// The multiReadCloser struct allows multiple readers to be concatenated
// as part of a single ReadCloser
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

// Extends strconv.ParseFloat to replace commas with blanks
func parseFloat(s string) float64 {
	result, err := strconv.ParseFloat(commaRegex.ReplaceAllString(s, ""), 64)
	if err != nil {
		panic(err)
	}

	return result
}
