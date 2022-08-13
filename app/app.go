package app

import (
	"github.com/andrewmolyuk/pixar"
	"github.com/andrewmolyuk/pixar/fileops"
	"github.com/andrewmolyuk/pixar/log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Pixar contains command line parameters and operating data
type Pixar struct {
	InputFolder    string `short:"i" long:"input" description:"Input folder" default:"."`
	OutputFolder   string `short:"o" long:"output" description:"Output folder" default:"output"`
	Move           bool   `short:"m" long:"move" description:"Move files instead of copying them"`
	Debug          bool   `short:"d" long:"debug" description:"Debug mode"`
	ShowVersion    bool   `short:"v" long:"version" description:"Show Pixar version info"`
	Extensions     string `short:"e" long:"extensions" description:"File extensions to process" default:".jpeg,.jpg,.tiff,.png"`
	BuildInfo      pixar.BuildInfo
	extensionsList []string // cache for extensions list
}

// Run is the main process where the application is running
func (p *Pixar) Run() {
	err := fileops.IsFolderExists(p.InputFolder)
	if err != nil {
		log.Error(err)
	}

	p.processFolder(p.InputFolder)
}

func (p *Pixar) processFolder(folder string) {
	log.Debug("Processing folder: \"%s\"", folder)
	files, err := os.ReadDir(folder)
	if err != nil {
		log.Error(err)
	}

	for _, f := range files {
		i, err := f.Info()
		if err != nil {
			log.Error(err)
		}

		if i.IsDir() {
			p.processFolder(folder + "/" + f.Name())
		} else {
			p.processFile(folder + "/" + f.Name())
		}
	}
}

func (p *Pixar) processFile(file string) {
	log.Debug("Processing file: \"%s\"", file)
	if p.isExtensionToProcess(file) {
		createDate, err := fileops.GetFileExifCreateDate(file)
		if err != nil {
			log.Warn("Error getting create date from file: \"%s\". Error: %s", file, err)
			return
		}
		p.processFileToOutput(file, createDate)
	}
}

func (p *Pixar) isExtensionToProcess(file string) bool {
	if p.extensionsList == nil {
		p.extensionsList = strings.Split(p.Extensions, ",")
	}

	extension := strings.ToLower(filepath.Ext(file))

	for _, e := range p.extensionsList {
		if extension == e {
			return true
		}
	}
	return false
}

func (p *Pixar) processFileToOutput(file string, date time.Time) {
	log.Debug("Processing file: \"%s\" to output", file)
	folder := p.OutputFolder + "/" + date.Format("2006/01/02")
	err := fileops.CreateFolder(folder)
	if err != nil {
		log.Error("Cannot create folder: \"%s\". Error: \"%s\"", folder, err)
	}
	if p.Move {
		err := fileops.MoveFile(file, folder)
		if err != nil {
			log.Error("Cannot move file: \"%s\" to folder: \"%s\". Error: \"%s\"", file, folder, err)
		}
	} else {
		err := fileops.CopyFile(file, folder)
		if err != nil {
			log.Error("Cannot copy file: \"%s\" to folder: \"%s\". Error: \"%s\"", file, folder, err)
		}
	}
}
