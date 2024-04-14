package benchmarks

import "sync"

func runQsBenchmark() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			var wg sync.WaitGroup
			wg.Add(5)
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 10, "q", "random")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 10, "q", "sorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 10, "q", "reverseSorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 10, "q", "thirdSorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 10, "q", "twoThirdsSorted")
			}()
			wg.Wait()
		}
	}
}
