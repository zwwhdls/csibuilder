package model

import "strings"

// safeImport returns a cleaned version of the provided string that can be used for imports
func safeImport(unsafe string) string {
	safe := unsafe

	// Remove dashes and dots
	safe = strings.Replace(safe, "-", "", -1)
	safe = strings.Replace(safe, ".", "", -1)

	return safe
}
