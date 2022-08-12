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
	InputFolder  string `short:"i" long:"input" description:"Input folder" default:"."`
	OutputFolder string `short:"o" long:"output" description:"Output folder" default:"output"`
	Move         bool   `short:"m" long:"move" description:"Move files instead of copying them"`
	Debug        bool   `short:"d" long:"debug" description:"Debug mode"`
	ShowVersion  bool   `short:"v" long:"version" description:"Show Pixar version info"`
	BuildInfo    pixar.BuildInfo
}

// DoWork is the main process where the application is getting started
func (a *Pixar) DoWork() {
	err := fileops.IsFolderExists(a.InputFolder)
	if err != nil {
		log.Error(err)
	}

	a.processFolder(a.InputFolder)
}

func (a *Pixar) processFolder(folder string) {
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
			a.processFolder(folder + "/" + f.Name())
		} else {
			a.processFile(folder + "/" + f.Name())
		}
	}
}

func (a *Pixar) processFile(file string) {
	log.Debug("Processing file: \"%s\"", file)
	if a.isImage(file) {
		createDate, err := fileops.GetFileExifCreateDate(file)
		if err != nil {
			log.Warn("Error getting create date from file: \"%s\". Error: %s", file, err)
			return
		}
		a.processFileToOutput(file, createDate)
	}
}

func (a *Pixar) isImage(file string) bool {
	extensions := map[string]bool{
		".jpeg": true,
		".jpg":  true,
		".tiff": true,
		".png":  true,
	}

	extension := strings.ToLower(filepath.Ext(file))

	if _, ok := extensions[extension]; ok {
		log.Debug("File: \"%s\" is an image file", file)
		return true
	}
	return false
}

func (a *Pixar) processFileToOutput(file string, date time.Time) {
	log.Debug("Processing file: \"%s\" to output", file)
	folder := a.OutputFolder + "/" + date.Format("2006/01/02")
	err := fileops.CreateFolder(folder)
	if err != nil {
		log.Error("Cannot create folder: \"%s\". Error: \"%s\"", folder, err)
	}
	if a.Move {
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
