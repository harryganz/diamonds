package main

import (
	"fmt"
	"os"

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
		fmt.Printf("File was %s. Num was %d\n", fileName, numDiamonds)
		return nil
	}

	app.Run(os.Args)
}
