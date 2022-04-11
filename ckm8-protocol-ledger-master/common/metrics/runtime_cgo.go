// +build cgo
// +build !appengine

ckm8 metrics

import "runtime"

func numCgoCall() int64 {
	return runtime.NumCgoCall()
}
