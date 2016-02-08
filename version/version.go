package version

import (
	"runtime"
)

var (
	Version   string = "0.1.0"
	Release   string = "a"
	GoVersion string = runtime.Version()
)
