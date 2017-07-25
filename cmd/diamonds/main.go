package main

import (
	"os"

	"github.com/harryganz/diamonds"
	"github.com/urfave/cli"
)

func main() {
	var fileName string
	var numDiamonds int
	var numThreads int

	app := cli.NewApp()
	app.Name = "diamonds"
	app.Version = "0.1.0"
	app.Usage = "Retrieves diamond info from diamond search engine and places it in csv file"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Value:       "diamonds.csv",
			Usage:       "Output file for data",
			Destination: &fileName,
		},
		cli.IntFlag{
			Name:        "num, n",
			Value:       20,
			Usage:       "Number of records to retrieve",
			Destination: &numDiamonds,
		},
		cli.IntFlag{
			Name: "threads, t",
			Value: 1,
			Usage: "Number of threads to use for querying search engine",
			Destination: &numThreads,
		},
	}

	app.Action = func(c *cli.Context) error {
		outputStream, err := os.Create(fileName)
		if err != nil {
			panic("Could not create file")
		}
		defer outputStream.Close()
		crawler := diamonds.NewCrawler(numDiamonds, outputStream)
		crawler.NumThreads = numThreads
		err = crawler.Crawl()

		return err
	}

	app.Run(os.Args)
}
