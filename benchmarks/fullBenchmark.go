package benchmarks

import (
	"log"
	"projekt1/timeTrack"
	"sync"
	"time"
)

func RunFullBenchmark() {
	log.Printf("Benchmark started")
	defer timeTrack.TimeTrack(time.Now(), "Benchmark")

	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		defer wg.Done()
		runQsBenchmark()
	}()
	go func() {
		defer wg.Done()
		runHsBenchmark()
	}()
	go func() {
		defer wg.Done()
		runSsBenchmark()
	}()
	go func() {
		defer wg.Done()
		runIsBenchmark()
	}()
	wg.Wait()
	log.Printf("Benchmark finished")
}
