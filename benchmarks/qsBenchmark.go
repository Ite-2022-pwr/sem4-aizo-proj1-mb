package benchmarks

func RunQsBenchmark() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 4; j++ {
			SingleConfigBenchmark(i, j, 0, 100, 0, 0)
			SingleConfigBenchmark(i, j, 0, 100, 0, 1)
			SingleConfigBenchmark(i, j, 0, 100, 0, 2)
			SingleConfigBenchmark(i, j, 0, 100, 0, 3)
			SingleConfigBenchmark(i, j, 0, 100, 0, 4)
		}
	}
}
