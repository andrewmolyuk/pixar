package main

import (
	"github.com/andrewmolyuk/pixar/app"
	"github.com/andrewmolyuk/pixar/exitor"
	"github.com/andrewmolyuk/pixar/log"
	"github.com/jessevdk/go-flags"

	"os"
)

func main() {
	a := app.Pixar{}

	parser := flags.NewParser(&a, flags.Default)
	parser.ShortDescription = "Pixar is command line pics archiver"
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
