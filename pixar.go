package pixar

// BuildInfo contains a minimal build information embedded into binary during version release
type BuildInfo struct {
	Version string
	Commit  string
}

// Action is a type of action to be performed by a command
type Action int

const (
	Copy Action = iota
	Move
	Skip
)

// FileAction is a type of action to be performed on file
type FileAction struct {
	File        string
	Destination string
	Action      Action
}
