package boundedPrimeSieve

import (
	"../run_bench/benchmark"
	"./goscr"
	"./base"
	"time"
)

var primesieveParams = []int{
	100, 1100, 2100, 3100, 4100, 5100, 6100, 7100, 8100, 9100,
}

func TimePrimeSieve(n int) time.Duration {
	start := time.Now()
	_ = goscr.PrimeSieve(n)
	elapsed := time.Since(start)
	return elapsed
}

func TimePrimeSieveBase(n int) time.Duration {
	start := time.Now()
	_ = base.PrimeSieve(n)
	elapsed := time.Since(start)
	return elapsed
}

func PrimeSieveBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieve)
	base_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieveBase)
	return scribble_results, base_results
}
