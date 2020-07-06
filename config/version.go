package config

import "github.com/lancer-kit/uwe/v2"

// The variables are set when compiling the binary and used to make sure which version of the backend is running.
// Example: go build -ldflags "-X lancer-kit/service-scaffold/config.version=$VERSION\
// -X lancer-kit/service-scaffold/config.build=$COMMIT \
// -X lancer-kit/service-scaffold/config.tag=$TAG" .

// nolint:gochecknoglobals
var (
	version = "n/a"
	build   = "n/a"
	tag     = "n/a"
)

func AppInfo() uwe.AppInfo {
	return uwe.AppInfo{
		Name:    ServiceName,
		Version: version,
		Build:   build,
		Tag:     tag,
	}
}
