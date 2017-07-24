package main

import (
	"os"

	"github.com/harryganz/diamonds"
	"github.com/urfave/cli"
)

func main() {
	var fileName string
	var numDiamonds int

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
			Value:       100,
			Usage:       "Number of records to retrieve",
			Destination: &numDiamonds,
		},
	}

	app.Action = func(c *cli.Context) error {
		outputStream := os.Stdout
		defer outputStream.Close()
		crawler := diamonds.NewCrawler(10, outputStream)
		err := crawler.Crawl()

		return err
	}

	app.Run(os.Args)
}
