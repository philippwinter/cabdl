package cabdl

import "github.com/cavaliercoder/grab"

type Context struct {
	Client *grab.Client
	URLFormat string
	Destination string
}
