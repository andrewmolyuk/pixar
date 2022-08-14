package pixar

// BuildInfo contains a minimal build information embedded into binary during version release
type BuildInfo struct {
	Version string
	Commit  string
}

// Action is a type of action to be performed by a command
type Action string

const (
	Copy Action = "Copy" // Copy is an action to copy a file
	Move Action = "Move" // Move is an action to move a file
	Skip Action = "Skip" // Skip is an action that skips the file
)

// FileAction is a type of action to be performed on file
type FileAction struct {
	File        string
	Destination string
	Action      Action
}
