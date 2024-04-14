package benchmarks

import "sync"

func runHsBenchmark() {
	for i := 0; i < 6; i++ {
		var wg = sync.WaitGroup{}
		wg.Add(5)
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "h", "random")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "h", "sorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "h", "reverseSorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "h", "thirdSorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "h", "twoThirdsSorted")
		}()
		wg.Wait()
	}
}
