package app

import (
	"encoding/csv"
	"fmt"
	"github.com/andrewmolyuk/pixar"
	"github.com/andrewmolyuk/pixar/log"
	"github.com/rwcarlsen/goexif/exif"
	"io"
	"os"
	"path/filepath"
	"time"
)

// MoveFile moves file to destination folder
func MoveFile(file string, folder string) error {
	log.Debug("Moving file: \"%s\" to folder: \"%s\"", file, folder)
	err := os.Rename(file, folder+"/"+filepath.Base(file))
	if err != nil {
		err = CopyFile(file, folder)
		if err != nil {
			return err
		}
		err = DeleteFile(file)
		if err != nil {
			return err
		}
	}

	log.Info("Moved file: \"%s\" to folder: \"%s\"", file, folder)
	return nil
}

// DeleteFile deletes file from disk
func DeleteFile(file string) error {
	err := os.RemoveAll(file)
	if err != nil {
		return err
	}
	return nil
}

// CopyFile copies file to destination folder
func CopyFile(file string, folder string) error {
	log.Debug("Copying file: \"%s\" to folder: \"%s\"", file, folder)

	src, err := os.Open(file)
	if err != nil {
		return err
	}
	defer func(file io.Closer) {
		err := file.Close()
		if err != nil {
			log.Warn("Error closing file: %s", file)
		}
	}(src)

	dst, err := os.Create(folder + "/" + filepath.Base(file))
	if err != nil {
		return err
	}
	defer func(file io.Closer) {
		err := file.Close()
		if err != nil {
			log.Warn("Error closing file: %s", file)
		}
	}(dst)

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	log.Info("Copied file: \"%s\" to folder: \"%s\"", file, folder)
	return nil
}

// CreateFolder creates folder if it doesn't exist
func CreateFolder(folder string) error {
	log.Debug("Creating folder: \"%s\"", folder)
	return os.MkdirAll(folder, 0755)
}

// IsFolderExists checks if folder exists
func IsFolderExists(folder string) error {
	f, err := os.Stat(folder)
	if err != nil {
		return err
	}

	if !f.IsDir() {
		return fmt.Errorf("folder: \"%s\" is not a folder", folder)
	}

	return nil
}

// GetFileExifCreateDate returns create date from file's EXIF information
func GetFileExifCreateDate(file string) time.Time {
	f, err := os.Open(file)
	defer func(file io.Closer) {
		err := file.Close()
		if err != nil {
			log.Error("Error closing file: %s", file)
		}
	}(f)

	if err != nil {
		return time.Time{}
	}

	exifData, err := exif.Decode(f)
	if err != nil {
		return time.Time{}
	}

	createDate, err := exifData.DateTime()
	if err != nil {
		return time.Time{}
	}

	return createDate
}

func WriteActionsToCsv(file string, actions []pixar.FileAction) error {
	log.Debug("Writing actions to CSV file: \"%s\"", file)
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func(file io.Closer) {
		err := file.Close()
		if err != nil {
			log.Error("Error closing file: %s", file)
		}
	}(f)

	writer := csv.NewWriter(f)
	defer writer.Flush()
	for _, a := range actions {
		err := writer.Write([]string{a.File, string(a.Action), a.Destination})
		if err != nil {
			return err
		}
	}
	return nil
}
