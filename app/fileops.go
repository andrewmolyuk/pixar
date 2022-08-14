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

func moveFile(file string, folder string) error {
	log.Debug("Moving file: \"%s\" to folder: \"%s\"", file, folder)
	err := os.Rename(file, folder+"/"+filepath.Base(file))
	if err != nil {
		err = copyFile(file, folder)
		if err != nil {
			return err
		}
		err = deleteFile(file)
		if err != nil {
			return err
		}
	}

	log.Info("Moved file: \"%s\" to folder: \"%s\"", file, folder)
	return nil
}

func deleteFile(file string) error {
	err := os.RemoveAll(file)
	if err != nil {
		return err
	}
	return nil
}

func copyFile(file string, folder string) error {
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

func createFolder(folder string) error {
	log.Debug("Creating folder: \"%s\"", folder)
	return os.MkdirAll(folder, 0755)
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

func writeActionsToCsv(file string, actions []pixar.FileAction) error {
	var zeroTime = time.Time{}
	log.Debug("Writing actions to CSV file: \"%s\"", file)

	err := createFolder(filepath.Dir(file))
	if err != nil {
		return err
	}

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
