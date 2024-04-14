package benchmarks

import "sync"

func runSsBenchmark() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			var wg sync.WaitGroup
			wg.Add(5)
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 10, "s", "random")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 10, "s", "sorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 10, "s", "reverseSorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 10, "s", "thirdSorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 10, "s", "twoThirdsSorted")
			}()
			wg.Wait()
		}
	}
}
