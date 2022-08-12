package app

import (
	"github.com/andrewmolyuk/pixar"
	"github.com/andrewmolyuk/pixar/log"
	"github.com/rwcarlsen/goexif/exif"
	"io"
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
	input, err := os.Stat(a.InputFolder)
	if err != nil {
		log.Error("Folder: \"%s\" does not exist", a.InputFolder)
	}

	if !input.IsDir() {
		log.Error("Folder: \"%s\" is not a folder", a.InputFolder)
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
		createDate, err := a.getFileExifCreateDate(file)
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

func (a *Pixar) getFileExifCreateDate(file string) (time.Time, error) {

	f, err := os.Open(file)
	defer func(file io.Closer) {
		err := file.Close()
		if err != nil {
			log.Error("Error closing file: %s", file)
		}
	}(f)

	if err != nil {
		log.Warn("Error opening file: \"%s\". Error: %s", file, err)
		return time.Time{}, err
	}

	exifData, err := exif.Decode(f)
	if err != nil {
		log.Warn("Error decoding file: \"%s\". Error: %s", file, err)
		return time.Time{}, err
	}

	createDate, err := exifData.DateTime()
	if err != nil {
		log.Warn("Error getting create date from file: \"%s\". Error: %s", file, err)
		return time.Time{}, err
	}

	return createDate, nil
}

func (a *Pixar) processFileToOutput(file string, date time.Time) {
	log.Debug("Processing file: \"%s\" to output", file)
	folder := a.OutputFolder + "/" + date.Format("2006/01/02")
	a.createFolder(folder)
	if a.Move {
		a.moveFile(file, folder)
	} else {
		a.copyFile(file, folder)
	}
}

func (a *Pixar) createFolder(folder string) {
	log.Debug("Creating folder: \"%s\"", folder)
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		log.Error("Error creating folder: \"%s\". Error: %s", folder, err)
	}
}

func (a *Pixar) moveFile(file string, folder string) {
	log.Debug("Moving file: \"%s\" to folder: \"%s\"", file, folder)
	err := os.Rename(file, folder+"/"+filepath.Base(file))
	if err != nil {
		log.Error("Error moving file: \"%s\" to folder: \"%s\". Error: %s", file, folder, err)
	}
	log.Info("Moved file: \"%s\" to folder: \"%s\"", file, folder)
}

func (a *Pixar) copyFile(file string, folder string) {
	log.Debug("Copying file: \"%s\" to folder: \"%s\"", file, folder)

	src, err := os.Open(file)
	if err != nil {
		log.Error("Error opening file: \"%s\". Error: %s", file, err)
		return
	}
	defer func(file io.Closer) {
		err := file.Close()
		if err != nil {
			log.Error("Error closing file: %s", file)
		}
	}(src)

	dst, err := os.Create(folder + "/" + filepath.Base(file))
	if err != nil {
		log.Error("Error creating file: \"%s\". Error: %s", folder+"/"+filepath.Base(file), err)
		return
	}
	defer func(file io.Closer) {
		err := file.Close()
		if err != nil {
			log.Error("Error closing file: %s", file)
		}
	}(dst)

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Error("Error copying file: \"%s\" to folder: \"%s\". Error: %s", file, folder, err)
	}

	log.Info("Copied file: \"%s\" to folder: \"%s\"", file, folder)
}
