package main

import (
	"gitlab.inn4science.com/gophers/service-scaffold/info"
)

var (
	Version = "0.2.0"
	Build   string
	Tag     string
)

func init() {
	info.App.Version = Version
	info.App.Build = Build
	info.App.Tag = Tag
}
