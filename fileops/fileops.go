package fileops

import (
	"github.com/andrewmolyuk/pixar/log"
	"io"
	"os"
	"path/filepath"
)

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

func DeleteFile(file string) error {
	err := os.RemoveAll(file)
	if err != nil {
		return err
	}
	return nil
}

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

func CreateFolder(folder string) error {
	log.Debug("Creating folder: \"%s\"", folder)
	return os.MkdirAll(folder, 0755)
}
