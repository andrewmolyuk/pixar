package log

// ILog is the simple logger interface
type ILog interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}
