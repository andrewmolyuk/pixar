package app

import (
	"github.com/andrewmolyuk/pixar"
	"github.com/andrewmolyuk/pixar/log"
	"github.com/andrewmolyuk/pixar/semaphore"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Pixar contains command line parameters and operating data
type Pixar struct {
	InputFolder      string             `short:"i" long:"input" description:"Input folder" default:"."`
	OutputFolder     string             `short:"o" long:"output" description:"Output folder" default:"output"`
	Move             bool               `short:"m" long:"move" description:"Move files instead of copying them"`
	Debug            bool               `short:"d" long:"debug" description:"Debug mode"`
	ShowVersion      bool               `short:"v" long:"version" description:"Show Pixar version info"`
	Extensions       string             `short:"e" long:"extensions" description:"File extensions to process" default:".jpeg,.jpg,.tiff,.png"`
	Simulation       bool               `short:"s" long:"simulation" description:"Simulation mode"`
	Csv              string             `short:"c" long:"csv" description:"CSV file name for actions output"`
	MaxConcurrentOps uint               `short:"n" long:"concurrent" description:"Maximum number of concurrent operations" default:"100"`
	DuplicatesPolicy string             `short:"p" long:"policy" description:"Policy for duplicates: skip, folder" default:"skip"`
	BuildInfo        pixar.BuildInfo    // BuildInfo contains a build info embedded into binary during version release
	extensionsList   []string           // cache for extensions list
	actions          []pixar.FileAction // actions to be performed on files
}

// Run is the main process where the application is running
func (p *Pixar) Run() {
	err := isFolderExists(p.InputFolder)
	if err != nil {
		log.Error(err)
	}

	p.processFolder(p.InputFolder)

	if p.Csv != "" {
		err := writeActionsToCsv(p.Csv, p.actions)
		if err != nil {
			log.Error(err)
		}
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
		action.ModificationDate = getFileModificationDate(file)
		exifCreateDate := getFileExifCreateDate(file)
		if exifCreateDate == zeroTime {
			action.Action = pixar.Skip
		} else {
			action.ExifCreateDate = exifCreateDate
			action.Destination = p.OutputFolder + "/" + exifCreateDate.Format("2006/01/02")
			if p.Move {
				action.Action = pixar.Move
			} else {
				action.Action = pixar.Copy
			}
		}
	} else {
		action.Action = pixar.Skip
	}
	action = p.detectDuplicates(action)
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
	s := semaphore.NewSemaphore(p.MaxConcurrentOps)
	for _, a := range p.actions {
		switch a.Action {
		case pixar.Copy:
			log.Info("Copying file: \"%s\" to \"%s\"", a.File, a.Destination)
			if !p.Simulation {
				s.Acquire()
				go copyFile(a.File, a.Destination, s)
			}

		case pixar.Move:
			log.Info("Moving file: \"%s\" to \"%s\"", a.File, a.Destination)
			if !p.Simulation {
				s.Acquire()
				go moveFile(a.File, a.Destination, s)
			}

		case pixar.Skip:
			log.Info("Skip file: \"%s\"", a.File)

		}
	}
	s.Wait()
}

func (p *Pixar) detectDuplicates(action pixar.FileAction) pixar.FileAction {
	log.Debug("Detecting duplicates for file: \"%s\"", action.File)

	if action.Action != pixar.Move && action.Action != pixar.Copy {
		return action
	}

	isDuplicate := false
	if _, err := os.Stat(action.Destination + "/" + filepath.Base(action.File)); err == nil {
		isDuplicate = true
	}
	if !isDuplicate {
		for _, a := range p.actions {
			if action.Destination+"/"+filepath.Base(action.File) == a.Destination+"/"+filepath.Base(a.File) {
				isDuplicate = true
				break
			}
		}
	}

	if isDuplicate {
		log.Warn("Duplicate file: \"%s\" in destination: \"%s\"", action.File, action.Destination)
		switch p.DuplicatesPolicy {
		case "skip":
			action.Action = pixar.Skip
		case "folder":
			action.Destination = p.OutputFolder + "/Duplicates/" + filepath.Dir(action.File)
		}
	}
	return action
}
