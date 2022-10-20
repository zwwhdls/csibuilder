package machinery

import (
	"fmt"
	"hash/fnv"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
)

// DefaultFuncMap returns the default template.FuncMap for rendering the template.
func DefaultFuncMap() template.FuncMap {
	return template.FuncMap{
		"title":      cases.Title,
		"lower":      strings.ToLower,
		"upper":      strings.ToUpper,
		"isEmptyStr": isEmptyString,
		"hashFNV":    hashFNV,
	}
}

// isEmptyString returns whether the string is empty
func isEmptyString(s string) bool {
	return s == ""
}

// hashFNV will generate a random string useful for generating a unique string
func hashFNV(s string) string {
	hasher := fnv.New32a()
	// Hash.Write never returns an error
	_, _ = hasher.Write([]byte(s))
	return fmt.Sprintf("%x", hasher.Sum(nil))
}
