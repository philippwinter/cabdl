package main

import (
	"github.com/urfave/cli"
	"time"
	"github.com/philippwinter/cabdl"
	"log"
	"os"
	"github.com/cavaliercoder/grab"
)

func main() {
	app := cli.NewApp()
	app.Name = "NY Cab Data"
	app.Version = "0.0.1"
	app.Action = entry
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "start_period",
			Value: "2009-01",
			Usage: "Download data starting in this period",
		},
		cli.StringFlag{
			Name: "end_period",
			Value: time.Now().AddDate(0, -1, 0).Format("2006-01"),
			Usage: "Download data ending in this period",
		},
		cli.StringFlag{
			Name: "destination",
			Value: ".",
			Usage: "Download folder, if not set defaults to current directory",
		},
		cli.StringFlag{
			Name: "url_format",
			Value: "https://s3.amazonaws.com/nyc-tlc/trip+data/yellow_tripdata_%s.csv",
			Usage: "Alternate download URL format",
		},
	}

	app.Run(os.Args)
}

func entry(context *cli.Context) {
	startPeriod, err := cabdl.AsPeriod(context.String("start_period"))
	if err != nil {
		log.Fatal(err)
	}

	endPeriod, err := cabdl.AsPeriod(context.String("end_period"))
	if err != nil {
		log.Fatal(err)
	}

	destination := context.String("destination")
	urlFormat := context.String("url_format")

	c := &cabdl.Context{
		Destination: destination,
		URLFormat: urlFormat,
		Client: grab.NewClient(),
	}

	cabdl.ForEachPeriod(*startPeriod, *endPeriod, c, cabdl.DownloadPeriodData)
}
