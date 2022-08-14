package pixar

import "time"

// BuildInfo contains a minimal build information embedded into binary during version release
type BuildInfo struct {
	Version string
	Commit  string
}

// Action is a type of action applied to file
type Action byte

// String returns a string representation of Action
func (s Action) String() string {
	switch s {
	case Copy:
		return "COPY"
	case Move:
		return "MOVE"
	case Skip:
		return "SKIP"
	}
	return "UNKNOWN"
}

// List of actions applied to files
const (
	Copy Action = iota
	Move
	Skip
)

// FileAction is a type of action to be performed on file
type FileAction struct {
	File             string
	Destination      string
	ModificationDate time.Time
	ExifCreateDate   time.Time
	Action           Action
}
