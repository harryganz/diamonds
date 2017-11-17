# Diamonds
Scraper for getting diamond price information

Scrapes diamond price information from the [diamond search engine database](http://diamondse.info).
The command line interface generates a CSV file from the scraped records.

This project is still in development.

### Requirements

* Go >= 1.8

### Installation

* Navigate to your Go workspace in the terminal (see the Go ["getting started" guide](https://golang.org/doc/install) for help 
on setting up a Go environment)
* Run ```go get github.com/harryganz/diamonds```
* Navigate to the diamonds directory ```cd {go_workspace}/src/github.com/harryganz/diamonds```
* Run ```go install```. This will install diamonds to your go binary directory ```{go_workspace}/bin```

### Usage

Example: Download 1000 diamond records to the current directory as a csv
```
./diamonds -n 1000 -f "diamonds.csv"
```

Usage options:
```
diamonds -n {number of records to retrieve} [-f {file name} -t {number of threads to use}]
```

