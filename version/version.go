package version

import (
	"fmt"
	"runtime"
)

var GitCommit string

var Environment string

var Version string

var BuildDate = ""

var GoVersion = runtime.Version()

var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
