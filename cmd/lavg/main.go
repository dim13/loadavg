package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/dim13/loadavg"
)

func main() {
	interval := flag.Duration("n", time.Second*5, "update interval")
	flag.Parse()
	ticker := time.NewTicker(*interval)
	defer ticker.Stop()
	for range ticker.C {
		fmt.Printf("%.2f\n", loadavg.Get())
	}
}
