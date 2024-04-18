package benchmarks

func RunHsBenchmark() {
	for i := 0; i < 6; i++ {
		SingleConfigBenchmark(i, 0, 0, 100, 1, 0)
		SingleConfigBenchmark(i, 0, 0, 100, 1, 1)
		SingleConfigBenchmark(i, 0, 0, 100, 1, 2)
		SingleConfigBenchmark(i, 0, 0, 100, 1, 3)
		SingleConfigBenchmark(i, 0, 0, 100, 1, 4)
	}
}
