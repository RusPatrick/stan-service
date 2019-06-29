package helpers

import (
	"runtime"
)

const (
	defaultDepth = 1
)

func GetFunctionName(depthArr ...int) string {
	var depth int
	if len(depthArr) == 0 {
		depth = defaultDepth
	} else {
		depth = depthArr[0]
	}
	pc, _, _, _ := runtime.Caller(depth)
	return runtime.FuncForPC(pc).Name()
}
