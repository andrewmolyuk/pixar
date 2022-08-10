package pixar

type BuildInfo struct {
	Version string
	Commit  string
}

type ILogger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
}
