package benchmarks

func runHsBenchmark() {
	for i := 0; i < 6; i++ {
		SingleConfigBenchmark(i, 0, 0, 100, "h", "random")
		SingleConfigBenchmark(i, 0, 0, 100, "h", "sorted")
		SingleConfigBenchmark(i, 0, 0, 100, "h", "reverseSorted")
		SingleConfigBenchmark(i, 0, 0, 100, "h", "thirdSorted")
		SingleConfigBenchmark(i, 0, 0, 100, "h", "twoThirdsSorted")
	}
}
