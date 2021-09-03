package gotestpkg

import "runtime"

func Version() string {
	return runtime.Version()
}
