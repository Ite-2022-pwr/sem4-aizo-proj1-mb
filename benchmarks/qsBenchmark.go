package benchmarks

func runQsBenchmark() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			SingleConfigBenchmark(i, j, 0, 10, "q", "random")
			SingleConfigBenchmark(i, j, 0, 10, "q", "sorted")
			SingleConfigBenchmark(i, j, 0, 10, "q", "reverseSorted")
			SingleConfigBenchmark(i, j, 0, 10, "q", "thirdSorted")
			SingleConfigBenchmark(i, j, 0, 10, "q", "twoThirdsSorted")
		}
	}
}
