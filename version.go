package main

import (
	"lancer-kit/service-scaffold/info"
)

// The variables are set when compiling the binary and used to make sure which version of the backend is running.
// Example: go build -ldflags "-X main.Version=$VERSION -X main.Build=$COMMIT -X main.Tag=$TAG" .
// nolint: gochecknoglobals
var (
	Version = "1.0.0-rc"
	Build   string
	Tag     string
)

// nolint: gochecknoinits
func init() {
	info.App.Version = Version
	info.App.Build = Build
	info.App.Tag = Tag
}
