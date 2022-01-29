package global

import (
	"fmt"
	"runtime"
)

const Version = "0.0.2-Beta"

func GetRuntime() string {
	return fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
