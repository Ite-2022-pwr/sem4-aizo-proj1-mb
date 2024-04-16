package benchmarks

import (
	"log"
	"projekt1/timeTrack"
	"time"
)

func RunFullBenchmark() {
	log.Printf("Benchmark started")
	defer timeTrack.TimeTrack(time.Now(), "Benchmark")
	runSsBenchmark()
	runHsBenchmark()
	runIsBenchmark()
	runQsBenchmark()
	log.Printf("Benchmark finished")
}
