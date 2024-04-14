package benchmarks

import "sync"

func runSsBenchmark() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			var wg = sync.WaitGroup{}
			wg.Add(5)
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 100, "s", "random")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 100, "s", "sorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 100, "s", "reverseSorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 100, "s", "thirdSorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, 0, j, 100, "s", "twoThirdsSorted")
			}()
			wg.Wait()
		}
	}
}
