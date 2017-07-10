package main

import (
  "log"
  "fmt"
  "net/url"
  "strconv"

  "github.com/PuerkitoBio/goquery"
)

func main () {
  // Create the URL object from the domain and path
  u, err := url.Parse("http://diamondse.info/webService.php")
  if err != nil {
    log.Fatal("Could not parse URL")
  }

  // Initialize a map of all the parameters using default values
  params := map[string]interface{} {
    "shape" : "Round",
    "minCarat" : 0.2,
    "maxCarat" : 30.0,
    "minColor" : 1,
    "maxColor" : 9,
    "minPrice" : 100.0,
    "maxPrice" : 1000000.0,
    "minCut" : 5,
    "maxCut" : 1,
    "minClarity" : 1,
    "maxClarity" : 10,
    "minDepth" : 0.0,
    "maxDepth" : 90.0,
    "minWidth" : 0.0,
    "maxWidth" : 90.0,
    "gia" : 1,
    "ags" : 1,
    "egl" : 0,
    "oth" : 0,
    "currency" : "USD",
    "rowStart" : 0,
    "sortCol": "price",
    "sortDir" : "ASC",
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

  // Request the URL and get the returned document
  doc, err := goquery.NewDocument(u.String())
  if err != nil {
    log.Fatal(err)
  }

  fmt.Print(doc.Text())
}
