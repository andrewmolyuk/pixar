package app

import (
	"github.com/andrewmolyuk/go-photos/log"
)

type GoPhotos struct {
	Debug bool `short:"d" long:"debug" description:"Debug mode"`
	Log   log.ILog
}

func (a *GoPhotos) DoWork() {
	a.Log.Debug("DoWork")
}
