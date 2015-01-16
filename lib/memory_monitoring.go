package lib

import (
	"log"
	"runtime"
	"time"
)

func StartMemoryMonitoring() {
	go func() {
		for {
			var stat runtime.MemStats
			runtime.ReadMemStats(&stat)
			log.Printf("%vKb", stat.Alloc/1024)
			time.Sleep(time.Second)
		}
	}()
}
