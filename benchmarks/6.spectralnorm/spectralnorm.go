package spectralnorm

import (
	"./base"
	"./goscr"
    "../run_bench/benchmark"
	"runtime"
	"time"
)

var nCPU = runtime.NumCPU()

var spectralNormParams = []int{
	100, 500, 1500, 2500, 3500, 4500, 5500,
}


func TimeSpectralNorm(n int) time.Duration {
	start := time.Now()
	goscr.SpectralNorm(n)
	elapsed := time.Since(start)
	return elapsed
}

func TimeSpectralNormBase(n int) time.Duration {
	start := time.Now()
	base.SpectralNorm(n)
	elapsed := time.Since(start)
	return elapsed
}

func SpectralNormBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(spectralNormParams, repetitions, TimeSpectralNorm)
	base_results := benchmark.TimeImpl(spectralNormParams, repetitions, TimeSpectralNormBase)
	return scribble_results, base_results
}
