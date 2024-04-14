package benchmarks

func runHsBenchmark() {
	for i := 0; i < 6; i++ {
		SingleConfigBenchmark(i, 0, 0, 10, "h", "random")
		SingleConfigBenchmark(i, 0, 0, 10, "h", "sorted")
		SingleConfigBenchmark(i, 0, 0, 10, "h", "reverseSorted")
		SingleConfigBenchmark(i, 0, 0, 10, "h", "thirdSorted")
		SingleConfigBenchmark(i, 0, 0, 10, "h", "twoThirdsSorted")
	}
}
