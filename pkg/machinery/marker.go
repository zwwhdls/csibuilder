package machinery

import (
	"fmt"
	"path/filepath"
	"strings"
)

const prefix = "+kubebuilder:scaffold:"

var commentsByExt = map[string]string{
	".go":   "//",
	".yaml": "#",
	".yml":  "#",
	// When adding additional file extensions, update also the NewMarkerFor documentation and error
}

// Marker represents a machine-readable comment that will be used for scaffolding purposes
type Marker struct {
	comment string
	value   string
}

// NewMarkerFor creates a new marker customized for the specific file
// Supported file extensions: .go, .yaml, .yml
func NewMarkerFor(path string, value string) Marker {
	ext := filepath.Ext(path)
	if comment, found := commentsByExt[ext]; found {
		return Marker{comment, value}
	}

	extensions := make([]string, 0, len(commentsByExt))
	for extension := range commentsByExt {
		extensions = append(extensions, fmt.Sprintf("%q", extension))
	}
	panic(fmt.Errorf("unknown file extension: '%s', expected one of: %s", ext, strings.Join(extensions, ", ")))
}

// String implements Stringer
func (m Marker) String() string {
	return m.comment + prefix + m.value
}

// EqualsLine compares a marker with a string representation to check if they are the same marker
func (m Marker) EqualsLine(line string) bool {
	line = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), m.comment))
	return line == prefix+m.value
}

// CodeFragments represents a set of code fragments
// A code fragment is a piece of code provided as a Go string, it may have multiple lines
type CodeFragments []string

// CodeFragmentsMap binds Markers and CodeFragments together
type CodeFragmentsMap map[Marker]CodeFragments
