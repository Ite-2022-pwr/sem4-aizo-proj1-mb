package benchmarks

func runSsBenchmark() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			SingleConfigBenchmark(i, 0, j, 10, "s", "random")
			SingleConfigBenchmark(i, 0, j, 10, "s", "sorted")
			SingleConfigBenchmark(i, 0, j, 10, "s", "reverseSorted")
			SingleConfigBenchmark(i, 0, j, 10, "s", "thirdSorted")
			SingleConfigBenchmark(i, 0, j, 10, "s", "twoThirdsSorted")
		}
	}
}
