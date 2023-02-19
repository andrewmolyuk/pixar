package app

import (
	"encoding/csv"
	"fmt"
	"github.com/andrewmolyuk/pixar"
	"github.com/andrewmolyuk/pixar/log"
	"github.com/andrewmolyuk/pixar/semaphore"
	"github.com/evanoberholster/imagemeta"
	"io"
	"os"
	"path/filepath"
	"time"
)

func moveFile(file string, folder string, s *semaphore.Semaphore) {
	log.Debug("Moving file: \"%s\" to folder: \"%s\"", file, folder)
	defer s.Release()

	failIfFileExists(file, folder)

	createFolder(folder)
	err := os.Rename(file, folder+"/"+filepath.Base(file))
	if err != nil {
		s.Acquire()
		copyFile(file, folder, s)
		err = deleteFile(file)
		if err != nil {
			log.Error("Error deleting file: %s. Error: %s", file, err)
		}
	}
}

func deleteFile(file string) error {
	err := os.RemoveAll(file)
	if err != nil {
		return err
	}
	return nil
}

func copyFile(file string, folder string, s *semaphore.Semaphore) {
	log.Debug("Copying file: \"%s\" to folder: \"%s\"", file, folder)
	defer s.Release()

	failIfFileExists(file, folder)

	createFolder(folder)
	src, err := os.Open(file)
	if err != nil {
		log.Error("Error opening file: %s. Error: %s", file, err)
	}
	defer closeFile(src)

	dst, err := os.Create(folder + "/" + filepath.Base(file))
	if err != nil {
		log.Error("Error creating file: %s. Error: %s", folder+"/"+filepath.Base(file), err)
	}
	defer closeFile(dst)

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Error("Error copying file: %s. Error: %s", file, err)
	}
}

func failIfFileExists(file string, folder string) {
	if _, err := os.Stat(folder + "/" + filepath.Base(file)); err == nil {
		log.Error("File: \"%s\" already exists in folder: \"%s\"", file, folder)
	}
}

func closeFile(file io.Closer) {
	err := file.Close()
	if err != nil {
		log.Error("Error closing file: %s. Error: %s", file, err)
	}
}

func createFolder(folder string) {
	log.Debug("Creating folder: \"%s\"", folder)
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		log.Error("Error creating folder: %s. Error:", folder, err)
	}
}

func isFolderExists(folder string) error {
	f, err := os.Stat(folder)
	if err != nil {
		return err
	}

	if !f.IsDir() {
		return fmt.Errorf("folder: \"%s\" is not a folder", folder)
	}

	return nil
}

func getFileModificationDate(file string) time.Time {
	f, err := os.Stat(file)
	if err != nil {
		return time.Time{}
	}

	return f.ModTime()
}

func getFileExifCreateDate(file string) time.Time {
	f, err := os.Open(file)
	defer closeFile(f)

	if err != nil {
		return time.Time{}
	}

	meta, err := imagemeta.Parse(f)
	if err != nil {
		return time.Time{}
	}

	exif, err := meta.Exif()
	if err != nil || exif == nil {
		return time.Time{}
	}

	createdDate, err := exif.DateTime(time.Local)
	if err != nil {
		return time.Time{}
	}

	return createdDate
}

func writeActionsToCsv(file string, actions []pixar.FileAction) error {
	var zeroTime = time.Time{}
	log.Debug("Writing actions to CSV file: \"%s\"", file)

	createFolder(filepath.Dir(file))

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

	err = writer.Write([]string{"file", "destination_folder", "modification_date", "exif_create_date", "action"})
	if err != nil {
		return err
	}

	for _, a := range actions {
		md := ""
		if a.ModificationDate != zeroTime {
			md = a.ModificationDate.String()
		}
		ecd := ""
		if a.ExifCreateDate != zeroTime {
			ecd = a.ExifCreateDate.String()
		}
		err := writer.Write([]string{a.File, a.Destination, md, ecd, a.Action.String()})
		if err != nil {
			return err
		}
	}
	return nil
}
