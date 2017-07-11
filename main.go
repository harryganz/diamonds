package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Create the URL object from the domain and path
	u, err := url.Parse("http://diamondse.info/webService.php")
	if err != nil {
		log.Fatal("Could not parse URL")
	}

	// Initialize a map of all the parameters using default values
	params := map[string]interface{}{
		"shape":      "Heart",
		"minCarat":   0.2,
		"maxCarat":   30.0,
		"minColor":   1,
		"maxColor":   9,
		"minPrice":   100.0,
		"maxPrice":   1000000.0,
		"minCut":     2,
		"maxCut":     1,
		"minClarity": 1,
		"maxClarity": 10,
		"minDepth":   0.0,
		"maxDepth":   90.0,
		"minWidth":   0.0,
		"maxWidth":   90.0,
		"gia":        1,
		"ags":        1,
		"egl":        0,
		"oth":        0,
		"currency":   "USD",
		"rowStart":   0,
		"sortCol":    "price",
		"sortDir":    "ASC",
	}

	// Create a query string from all of the params
	q := u.Query()
	var queryParam string
	for k, v := range params {
		// Convert param from int/float/string to string
		switch val := v.(type) {
		case string:
			queryParam = val
		case int:
			queryParam = strconv.Itoa(val)
		case float64:
			queryParam = strconv.FormatFloat(val, 'f', 3, 64)
		default:
			panic("Unrecognized type for parameter")
		}

		q.Set(k, queryParam)
	}
	// Append the query to the URL
	u.RawQuery = q.Encode()

	// Request the URL and get the response
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	// Appends html, body, and table tags around returned data because
	// parser cannot handle tr tags not in a table
	// TODO use a writer to append instead so that this can be done
	// asynchronously
	body, err := ioutil.ReadAll(resp.Body)
	if err = resp.Body.Close(); err != nil {
		log.Fatal(err)
	}
	body = append([]byte("<html><body><table><tbody>"), body...)
	body = append(body, []byte("</tbody></table><body></html>")...)

	// Parse document with goquery
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%10s %10s %10s %10s %10s %10s %10s %10s %20s %10s\n",
		"Shape", "Carat", "Cut", "Color", "Clarity",
		"Table", "Depth", "Cert", "Dimensions", "Price")
	doc.Find("tr").Each(func(i int, row *goquery.Selection) {
		diamond := make(map[string]string)
		row.Find("td").Each(func(k int, cell *goquery.Selection) {
			switch k {
			case 1:
				diamond["shape"] = cell.Text()
			case 2:
				diamond["carat"] = cell.Text()
			case 3:
				diamond["cut"] = cell.Text()
			case 4:
				diamond["color"] = cell.Text()
			case 5:
				diamond["clarity"] = cell.Text()
			case 6:
				diamond["table"] = cell.Text()
			case 7:
				diamond["depth"] = cell.Text()
			case 8:
				diamond["cert"] = cell.Text()
			case 9:
				diamond["dimensions"] = cell.Text()
			case 10:
				diamond["price"] = strings.Replace(cell.Text(), "$", "", 1)
			}
		})

		fmt.Printf("%10s %10s %10s %10s %10s %10s %10s %10s %20s %10s\n",
			diamond["shape"], diamond["carat"], diamond["cut"], diamond["color"],
			diamond["clarity"], diamond["table"], diamond["depth"], diamond["cert"],
			diamond["dimensions"], diamond["price"])
	})
}
