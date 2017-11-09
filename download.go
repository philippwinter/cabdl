package cabdl

import (
	"fmt"
	"os"
	"time"

	"github.com/cavaliercoder/grab"
	"log"
)

func DownloadPeriodData(context *Context, period Period) {
	url := fmt.Sprintf(context.URLFormat, period)

	download(context.Client, url, context.Destination)
}

func download(client *grab.Client, url string, destination string) {
	req, _ := grab.NewRequest(destination, url)

	// start download
	log.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	log.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Download failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)
}
