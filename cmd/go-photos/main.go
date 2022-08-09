package main

import (
	"github.com/andrewmolyuk/go-photos/app"
	"github.com/andrewmolyuk/go-photos/exitor"
	"github.com/andrewmolyuk/go-photos/log"
	"github.com/jessevdk/go-flags"

	"os"
)

func main() {
	a := app.GoPhotos{}

	parser := flags.NewParser(&a, flags.Default)
	parser.ShortDescription = "go-photos command line application"
	parser.LongDescription = "Scan folder and move photos into folders according their EXIF information."

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	a.Log = log.New(a.Debug, nil, true, exitor.Default())

	a.DoWork()
}
