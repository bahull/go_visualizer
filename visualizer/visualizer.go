package visualizer

import (
	"fmt"
	"github.com/bahull/go_visualizer/telemetry"
)

func PrintStats(stats telemetry.RuntimeStats) {
	fmt.Printf("[%s] Goroutines: %d | Alloc: %d KB | Sys: %d KB | GC: %d\n",
		stats.Timestamp.Format("15:04:05"),
		stats.NumGoroutine,
		stats.Alloc/1024,
		stats.Sys/1024,
		stats.NumGC,
	)
}
