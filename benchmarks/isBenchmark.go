package benchmarks

import "sync"

func runIsBenchmark() {
	var wg = sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		wg.Add(5)
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "i", "random")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "i", "sorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "i", "reverseSorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "i", "thirdSorted")
		}()
		go func() {
			defer wg.Done()
			SingleConfigBenchmark(i, 0, 0, 100, "i", "twoThirdsSorted")
		}()
	}
	wg.Wait()
}
