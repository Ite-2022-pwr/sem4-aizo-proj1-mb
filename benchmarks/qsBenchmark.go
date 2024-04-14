package benchmarks

func runQsBenchmark() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			SingleConfigBenchmark(i, j, 0, 100, "q", "random")
			SingleConfigBenchmark(i, j, 0, 100, "q", "sorted")
			SingleConfigBenchmark(i, j, 0, 100, "q", "reverseSorted")
			SingleConfigBenchmark(i, j, 0, 100, "q", "thirdSorted")
			SingleConfigBenchmark(i, j, 0, 100, "q", "twoThirdsSorted")
		}
	}
}
