package benchmarks

func RunIsBenchmark() {
	for i := 0; i < 6; i++ {
		SingleConfigBenchmark(i, 0, 0, 100, 2, 0)
		SingleConfigBenchmark(i, 0, 0, 100, 2, 1)
		SingleConfigBenchmark(i, 0, 0, 100, 2, 2)
		SingleConfigBenchmark(i, 0, 0, 100, 2, 3)
		SingleConfigBenchmark(i, 0, 0, 100, 2, 4)
	}
}
