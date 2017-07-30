package skel

import (
	"fmt"
	"github.com/hashicorp/go-version"
)

// Version The main, base version of the project.
const Version = "0.1.0"

// VersionPrerelease A pre-release marker for version.
const VersionPrerelease = "dev"

// GitCommit The git commit of the build; this is set by the compiler.
var GitCommit string

// SemVersion An instance of the version object; this can be used during compile
// and init time to // ensure that version of application is proper
// (matching semverse.org).
var SemVersion = version.Must(version.NewVersion(Version))

// VersionString returns the semantic version of skel.
func VersionString() string {
	if VersionPrerelease != "" {
		return fmt.Sprintf("%s-%s", Version, VersionPrerelease)
	}
	return Version
}
