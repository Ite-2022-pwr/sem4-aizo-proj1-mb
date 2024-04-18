package benchmarks

func RunSsBenchmark() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			SingleConfigBenchmark(i, 0, j, 100, 3, 0)
			SingleConfigBenchmark(i, 0, j, 100, 3, 1)
			SingleConfigBenchmark(i, 0, j, 100, 3, 2)
			SingleConfigBenchmark(i, 0, j, 100, 3, 3)
			SingleConfigBenchmark(i, 0, j, 100, 3, 4)
		}
	}
}
