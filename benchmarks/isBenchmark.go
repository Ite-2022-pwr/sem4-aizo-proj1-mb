package benchmarks

func runIsBenchmark() {
	for i := 0; i < 6; i++ {
		SingleConfigBenchmark(i, 0, 0, 10, "i", "random")
		SingleConfigBenchmark(i, 0, 0, 10, "i", "sorted")
		SingleConfigBenchmark(i, 0, 0, 10, "i", "reverseSorted")
		SingleConfigBenchmark(i, 0, 0, 10, "i", "thirdSorted")
		SingleConfigBenchmark(i, 0, 0, 10, "i", "twoThirdsSorted")
	}
}
