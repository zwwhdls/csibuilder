{{ .Boilerplate }}

package csi

import (
	"encoding/json"
	"fmt"
	"runtime"
)

// These are set during build time via -ldflags
var (
	driverVersion string
	gitCommit     string
	buildDate     string
)

// VersionInfo struct
type VersionInfo struct {
	DriverVersion string
	GitCommit     string
	BuildDate     string
	GoVersion     string
	Compiler      string
	Platform      string
}

// GetVersion returns VersionInfo
func GetVersion() VersionInfo {
	return VersionInfo{
		DriverVersion: driverVersion,
		GitCommit:     gitCommit,
		BuildDate:     buildDate,
		GoVersion:     runtime.Version(),
		Compiler:      runtime.Compiler,
		Platform:      fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

// GetVersionJSON returns version in JSON
func GetVersionJSON() (string, error) {
	info := GetVersion()
	marshalled, err := json.MarshalIndent(&info, "", "  ")
	if err != nil {
		return "", err
	}
	return string(marshalled), nil
}
