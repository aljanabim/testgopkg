package testgopkg

import "runtime"

func Version() string {
	return runtime.Version()
}
