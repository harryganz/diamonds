package diamonds

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the url for the diamonds search engine database
const DefaultBaseURL = "http://diamondse.info/webService.php"

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

// DefaultPageGetter returns a page from the diamond search engine using the defaults
func DefaultPageGetter(baseURL string, params Parameters, rowStart int) (io.Reader, error) {
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