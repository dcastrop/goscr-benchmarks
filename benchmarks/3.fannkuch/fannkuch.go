package fannkuch

import (
    "../run_bench/benchmark"
    "./base"
    "./goscr"
	"time"
)

const NCHUNKS = 720

var fannkuchParams = []int{
   4, 
   5,
   6,
   7,
   8,
   9,
   10,
   11,
   12,
}

func TimeFannkuch(n int) time.Duration {
	start := time.Now()
	goscr.Fannkuch(n)
	elapsed := time.Since(start)
	return elapsed
}

func TimeFannkuchBase(n int) time.Duration {
	start := time.Now()
	base.Fannkuch(n)
	return time.Since(start)
}

func FannkuchBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(fannkuchParams, repetitions, TimeFannkuch)
	base_results := benchmark.TimeImpl(fannkuchParams, repetitions, TimeFannkuchBase)
	return scribble_results, base_results
}
