package exitor

import "os"

// IExitor is the simple interface for os.Exit
type IExitor interface {
	Exit(int)
}

// Ensure defaultExitor implements IExitor interface
var _ IExitor = (*defaultExitor)(nil)

type defaultExitor struct{}

func (e *defaultExitor) Exit(code int) {
	os.Exit(code)
}

// Default creates a new instance of exitor
func Default() IExitor {
	return &defaultExitor{}
}
