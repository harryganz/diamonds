package main

import (
	"fmt"
	"bytes"
	"strconv"
	"net/url"
	"net/http"
	"io"
	"io/ioutil"
	"os"

	"github.com/harryganz/diamonds"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
)

const (
	baseUrl = "http://diamondse.info/webService.php"
)

func nopReaderCloser() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBuffer([]byte{}))
}

type multiReadCloser struct {
	reader io.Reader
	close func () error
}

func (mrc multiReadCloser) Read(p []byte) (int, error) {
	return mrc.reader.Read(p)
}

func (mrc multiReadCloser) Close () error {
	return mrc.close()
}

func getPage(params map[string]string, rowStart int) (io.ReadCloser, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nopReaderCloser(), err
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}

	q.Set("rowStart", strconv.Itoa(rowStart))

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
		close: resp.Body.Close,
	}

	return mrc, nil
}

func parseFloat (s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

func parsePage(page io.Reader) ([]diamonds.Diamond, error) {
	results := []diamonds.Diamond{}
	doc, err := goquery.NewDocumentFromReader(page)
	if err != nil {
		return []diamonds.Diamond{}, err
	}

	doc.Find("tr").Each(func (i int, s *goquery.Selection) {
		diamond := diamonds.Diamond{}
		s.Find("td").Each(func(k int, ss *goquery.Selection) {
			text := ss.Text()

			switch k {
			case 1:
				diamond.Shape = text
			case 2:
				diamond.Carat, _ = parseFloat(text)
			case 3:
				diamond.Cut = text
			case 4:
				diamond.Color = text
			case 5:
				diamond.Clarity = text
			case 6:
				diamond.Width, _ = parseFloat(text)
			case 7:
				diamond.Depth, _ = parseFloat(text)
			case 8:
				diamond.Certification = text
			case 9:
				diamond.Dimensions = text
			case 10:
				diamond.Price, _ = parseFloat(text[1:])
			}
		})

		results = append(results, diamond)
	})

	return results, nil
}

func main() {
	params := map[string]string {
		"shape": "none",
		"minCarat": "0.20",
		"maxCarat": "30.00",
		"minColor": "1",
		"maxColor": "9",
		"minPrice": "100",
		"maxPrice": "1000000",
		"minCut": "5",
		"maxCut": "1",
		"minClarity": "1",
		"maxClarity": "10",
		"minDepth": "0.00",
		"maxDepth": "90.00",
		"minWidth": "0.00",
		"maxWidth": "90.00",
		"gia": "1",
		"ags": "1",
		"egl": "0",
		"oth": "0",
		"currency": "USD",
		"sortCol": "price",
		"sortDir": "ASC",
	}

	results := []diamonds.Diamond{}
	for i := 0; i < 5000; i += 20 {
		page, err := getPage(params, i)
		defer page.Close()
		if err != nil {
			fmt.Println(err)
			return
		}

		pageResults, err := parsePage(page)
		if err != nil {
			fmt.Println(err)
			return
		}
		results = append(results, pageResults...)
	}

	outFile, err := os.Create("diamonds.csv")
	defer outFile.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = gocsv.MarshalFile(results, outFile)
	if err != nil {
		fmt.Println(err)
		return
	}
}
