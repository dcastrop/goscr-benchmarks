package benchmark

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type BenchmarkTimes map[int][]time.Duration

type TimeFunction func(int) time.Duration

const (
	SECOND = 1000000000
)

var MinExecTime = 10000000000

func TimeImpl(values []int, repetitions int, implFunction TimeFunction) BenchmarkTimes {
	results := make(map[int][]time.Duration)
	for _, val := range values {
		rs := NewRunningStat()
		// times := make([]time.Duration, repetitions)
		var times []time.Duration
		// for j := 0; j < repetitions; j++ {
		var j int
		requirement := 0.0
		for j = 0; rs.Mean()*float64(j) < float64(MinExecTime) || j < repetitions || rs.StandardDeviation()/rs.Mean() > 0.05+requirement; j++ {
			times = append(times, implFunction(val))
			rs.Push(float64(times[j]))
			requirement += 0.00001
		}
		fmt.Printf("val: %d - repetitions: %d, mean: %f, sd: %f\n", val, j, rs.Mean(), rs.StandardDeviation())
		results[val] = times
	}
	return results
}

func Sum(times []time.Duration) int64 {
	sum := int64(0)
	for _, duration := range times {
		sum += duration.Microseconds()
	}
	return sum
}

func Average(times []time.Duration) int64 {
	sum := Sum(times)
	return sum / int64(len(times))
}

func StandardDeviation(times []time.Duration) float64 {
	numValues := float64(len(times))
	avg := float64(Sum(times)) / numValues
	sd := float64(0)
	for _, timing := range times {
		sd += math.Pow(float64(timing.Microseconds())-avg, 2)
	}
	sd = math.Sqrt(sd / numValues)
	return sd
}

func ResultsToString(name string, results BenchmarkTimes) string {
	output := make([]string, len(results))
	idx := 0
	for key, times := range results {
		avg := Average(times)
		sd := StandardDeviation(times)
		header := fmt.Sprintf("%d - mean: %d, sd: %f", key, avg, sd)

		strResults := make([]string, len(times))
		for i, timing := range times {
			strResults[i] = strconv.Itoa(int(timing.Microseconds()))
		}
		output[idx] = fmt.Sprintf("%s\n%s", header, strings.Join(strResults, " "))
		idx++
	}
	return fmt.Sprintf("%s\n%s", name, strings.Join(output, "\n"))
}
