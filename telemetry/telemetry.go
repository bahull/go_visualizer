package telemetry

import (
	"runtime"
	"time"
)

type RuntimeStats struct {
	Timestamp       time.Time
	NumGoroutine    int
	Alloc           uint64
	TotalAlloc      uint64
	Sys             uint64
	NumGC           uint32
}

func Collect() RuntimeStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return RuntimeStats{
		Timestamp:    time.Now(),
		NumGoroutine: runtime.NumGoroutine(),
		Alloc:        m.Alloc,
		TotalAlloc:   m.TotalAlloc,
		Sys:          m.Sys,
		NumGC:        m.NumGC,
	}
}