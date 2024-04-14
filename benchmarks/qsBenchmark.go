package benchmarks

import "sync"

func runQsBenchmark() {
	var wg = sync.WaitGroup{}
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			wg.Add(5)
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 100, "q", "random")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 100, "q", "sorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 100, "q", "reverseSorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 100, "q", "thirdSorted")
			}()
			go func() {
				defer wg.Done()
				SingleConfigBenchmark(i, j, 0, 100, "q", "twoThirdsSorted")
			}()
		}
	}
	wg.Wait()
}
