package pixar

// BuildInfo contains a minimal build information embedded into binary during version release
type BuildInfo struct {
	Version string
	Commit  string
}
