package model

import "strings"

// Resource contains the information required to scaffold files for a resource.
type Resource struct {
	// Path is the path to the go package where the types are defined.
	Path           string `json:"path,omitempty"`
	CSIName        string `json:"csiName"`
	AttachRequired bool   `json:"attachRequired,omitempty"`
}

func (r Resource) Replacer() *strings.Replacer {
	var replacements []string

	return strings.NewReplacer(replacements...)
}
