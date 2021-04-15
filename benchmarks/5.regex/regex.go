package regex

import (
    "../run_bench/benchmark"
    "./goscr"
    "./base"
	"fmt"
	"io/ioutil"
	"os"
	"time"

)

var regexreduxParams = []int{
	0, 1, 2, 3, 4, 5, 6, 7,
}

var regexreduxFiles = []string{
	"regexredux-input1000.txt",
	"regexredux-input10000.txt",
	"regexredux-input100000.txt",
	"regexredux-input500000.txt",
	"regexredux-input1000000.txt",
	"regexredux-input5000000.txt",
	"regexredux-input10000000.txt",
	"regexredux-input25000000.txt",
}

func TimeRegexRedux(n int) time.Duration {
	b := readFile(regexreduxFiles[n])
	start := time.Now()
	goscr.RegexRedux(b)
	elapsed := time.Since(start)
	return elapsed
}

func TimeRegexReduxBase(n int) time.Duration {
	b := readFile(regexreduxFiles[n])
	start := time.Now()
	base.RegexRedux(b)
	elapsed := time.Since(start)
	return elapsed
}

func readFile(file string) []byte {
	f, err := os.Open(fmt.Sprintf("./data/%s", file))
	if err != nil {
		panic("Can't open input file")
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Can't read input file")
	}
	err = f.Close()
	if err != nil {
		panic("Can't close input file")
	}
	return b
}

func RegexReduxBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(regexreduxParams, repetitions, TimeRegexRedux)
	base_results := benchmark.TimeImpl(regexreduxParams, repetitions, TimeRegexReduxBase)
	return scribble_results, base_results
}
