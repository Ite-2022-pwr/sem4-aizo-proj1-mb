package benchmarks

import "sync"

func RunFullBenchmark() {

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
}