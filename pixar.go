package pixar

// BuildInfo contains a minimal build information embedded into binary during version release
type BuildInfo struct {
	Version string
	Commit  string
}

// Action is a type of action applied to file
type Action string

// List of actions applied to files
const (
	Copy Action = "Copy"
	Move Action = "Move"
	Skip Action = "Skip"
)

// FileAction is a type of action to be performed on file
type FileAction struct {
	File        string
	Destination string
	Action      Action
}
