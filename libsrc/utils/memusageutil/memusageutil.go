package memusageutil

import (
	"runtime"
)

func LogMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// trace.Lg("Alloc<%v:bytes,%v:mb> TotalAlloc<%v:bytes,%v:mb> Sys<%v:bytes,%v:mb> NumGC<%v>",m.Alloc, m.Alloc/1024/1024, m.TotalAlloc, m.TotalAlloc/1024/1024, m.Sys, m.Sys/1024/1024, m.NumGC)
}
