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
	InputFolder    string             `short:"i" long:"input" description:"Input folder" default:"."`
	OutputFolder   string             `short:"o" long:"output" description:"Output folder" default:"output"`
	Move           bool               `short:"m" long:"move" description:"Move files instead of copying them"`
	Debug          bool               `short:"d" long:"debug" description:"Debug mode"`
	ShowVersion    bool               `short:"v" long:"version" description:"Show Pixar version info"`
	Extensions     string             `short:"e" long:"extensions" description:"File extensions to process" default:".jpeg,.jpg,.tiff,.png"`
	Simulation     bool               `short:"s" long:"simulation" description:"Simulation mode"`
	Csv            bool               `short:"c" long:"csv" description:"Output to CSV file"`
	BuildInfo      pixar.BuildInfo    // BuildInfo contains a build info embedded into binary during version release
	extensionsList []string           // cache for extensions list
	actions        []pixar.FileAction // actions to be performed on files
}

// Run is the main process where the application is running
func (p *Pixar) Run() {
	err := fileops.IsFolderExists(p.InputFolder)
	if err != nil {
		log.Error(err)
	}

	p.processFolder(p.InputFolder)

	if p.Simulation {
		log.Warn("SIMULATION MODE")
	}
	p.performActions()
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
			p.defineFileAction(folder + "/" + f.Name())
		}
	}
}

func (p *Pixar) defineFileAction(file string) {
	var zeroTime = time.Time{}
	log.Debug("Processing file: \"%s\"", file)
	action := pixar.FileAction{
		File: file,
	}
	if p.isExtensionToProcess(file) {
		createDate := fileops.GetFileExifCreateDate(file)
		if createDate == zeroTime {
			action.Action = pixar.Skip
		} else {

			action.Destination = p.OutputFolder + "/" + createDate.Format("2006/01/02")
			if p.Move {
				action.Action = pixar.Move
			} else {
				action.Action = pixar.Copy
			}
		}
	} else {
		action.Action = pixar.Skip
	}
	p.actions = append(p.actions, action)
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

func (p *Pixar) performActions() {
	for _, a := range p.actions {
		switch a.Action {
		case pixar.Copy:
			log.Info("Copying file: \"%s\" to \"%s\"", a.File, a.Destination)
			if !p.Simulation {
				err := fileops.CreateFolder(a.Destination)
				if err != nil {
					log.Error("Cannot create folder: \"%s\". Error: \"%s\"", a.Destination, err)
				}
				err = fileops.CopyFile(a.File, a.Destination)
				if err != nil {
					log.Error("Cannot copy file: \"%s\" to folder: \"%s\". Error: \"%s\"", a.File, a.Destination, err)
				}
			}
		case pixar.Move:
			log.Info("Moving file: \"%s\" to \"%s\"", a.File, a.Destination)
			if !p.Simulation {
				err := fileops.CreateFolder(a.Destination)
				if err != nil {
					log.Error("Cannot create folder: \"%s\". Error: \"%s\"", a.Destination, err)
				}
				err = fileops.MoveFile(a.File, a.Destination)
				if err != nil {
					log.Error("Cannot move file: \"%s\" to folder: \"%s\". Error: \"%s\"", a.File, a.Destination, err)
				}
			}
		case pixar.Skip:
			log.Info("Skip file: \"%s\"", a.File)

		}
	}
}
