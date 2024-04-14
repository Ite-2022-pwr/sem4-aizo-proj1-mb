package benchmarks

import "sync"

func runIsBenchmark() {
	for i := 0; i < 6; i++ {
		var wg sync.WaitGroup
		wg.Add(5)
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 10, "i", "random")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 10, "i", "sorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 10, "i", "reverseSorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 10, "i", "thirdSorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 10, "i", "twoThirdsSorted")
		}()
		wg.Wait()
	}
}
