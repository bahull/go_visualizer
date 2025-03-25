package main

import (
	"fmt"
	"time"
	"github.com/bahull/go_visualizer/telemetry"
	"github.com/bahull/go_visualizer/visualizer"
)

func main() {
	interval := time.Nanosecond
	fmt.Println("Starting Go Runtime Visualizer (CLI Mode)")

	for {
		stats := telemetry.Collect()
		visualizer.PrintStats(stats)
		time.Sleep(interval)
	}
}
