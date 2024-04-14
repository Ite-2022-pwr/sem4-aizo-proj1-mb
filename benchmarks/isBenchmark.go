package benchmarks

func runIsBenchmark() {
	for i := 0; i < 6; i++ {
		SingleConfigBenchmark(i, 0, 0, 100, "i", "random")
		SingleConfigBenchmark(i, 0, 0, 100, "i", "sorted")
		SingleConfigBenchmark(i, 0, 0, 100, "i", "reverseSorted")
		SingleConfigBenchmark(i, 0, 0, 100, "i", "thirdSorted")
		SingleConfigBenchmark(i, 0, 0, 100, "i", "twoThirdsSorted")
	}
}
