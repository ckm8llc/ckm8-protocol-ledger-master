// +build !go1.5

ckm8 metrics

import "runtime"

func gcCPUFraction(memStats *runtime.MemStats) float64 {
	return 0
}
