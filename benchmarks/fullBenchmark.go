package benchmarks

import (
	"log"
	"projekt1/timeTrack"
	"time"
)

func RunFullBenchmark() {
	log.Printf("Benchmark started")
	defer timeTrack.TimeTrack(time.Now(), "Benchmark")
	RunSsBenchmark()
	RunHsBenchmark()
	RunIsBenchmark()
	RunQsBenchmark()
	log.Printf("Benchmark finished")
}
