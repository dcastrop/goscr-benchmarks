package boundedfib

import (
	"./base"
	"./goscr"
    "../run_bench/benchmark"
	"time"
)

var boundedfibParams = []int{
	5, 10, 15, 20, 25,
}

func TimeBoundedFibonacci(n int) time.Duration {
	start := time.Now()
	_ = goscr.Fibonacci(n)
	return time.Since(start)
}

func TimeBoundedFibonacciBase(n int) time.Duration {
	start := time.Now()
	_ = base.Fibonacci(n)
	return time.Since(start)
}

func BoundedFibonacciBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacci)
	base_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacciBase)
	return scribble_results, base_results
}
