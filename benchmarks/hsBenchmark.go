package benchmarks

import "sync"

func runHsBenchmark() {
	var wg = sync.WaitGroup{}
	for i := 0; i < 6; i++ {
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
	}
	wg.Wait()
}
