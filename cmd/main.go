package main

import (
	"fmt"
	"github.com/andrewmolyuk/pixar"
	"github.com/andrewmolyuk/pixar/app"
	"github.com/andrewmolyuk/pixar/exitor"
	"github.com/andrewmolyuk/pixar/log"
	"github.com/jessevdk/go-flags"

	"os"
)

var (
	version = "0.1.2.DEVELOPMENT"
	commit  = "UNKNOWN"
)

func main() {
	pxr := app.Pixar{
		BuildInfo: pixar.BuildInfo{
			Version: version,
			Commit:  commit,
		},
	}

	parser := flags.NewParser(&pxr, flags.Default)
	parser.ShortDescription = "Pixar is command line pics archiver"
	parser.LongDescription = "Scan folders and move photos and videos into folders according to their EXIF information"

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	if pxr.ShowVersion {
		fmt.Printf("Pixar %s\n", fmt.Sprintf("%s (git: %s)", version, commit[:7]))
		os.Exit(0)
	}

	if pxr.InputFolder == "" {
		fmt.Println("Input folder is not specified")
		os.Exit(1)
	}

	pxr.Log = log.New(pxr.Debug, nil, true, exitor.Default())

	pxr.DoWork()
}
