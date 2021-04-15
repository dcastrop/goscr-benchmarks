package knuc 

import (
    "../run_bench/benchmark"
    "./base"
    "./goscr"
    "os"
    "fmt"
    "io/ioutil"
	"bufio"
	"bytes"
	"time"
)

var kNucleotideParams = []int{
	0, 1, 2, 3, 4, 5, 6, 7,
}
var knucleotideFiles = []string{
	"knucleotide-input1000.txt",
	"knucleotide-input10000.txt",
	"knucleotide-input50000.txt",
	"knucleotide-input100000.txt",
	"knucleotide-input500000.txt",
	"knucleotide-input1000000.txt",
	"knucleotide-input2500000.txt",
	"knucleotide-input5000000.txt",
}

func NewKNucleotideEnv(n int) []byte {
	b := readFile(knucleotideFiles[n])
	dna := toBits(readSequence(">THREE", b))
	return dna
}

func TimeKNucleotide(n int) time.Duration {
	dna := NewKNucleotideEnv(n)
	start := time.Now()
	goscr.KNucleotide(dna)
	elapsed := time.Since(start)
	return elapsed
}

func TimeKNucleotideBase(n int) time.Duration {
	dna := NewKNucleotideEnv(n)
	start := time.Now()
	base.KNucleotide(dna)
	elapsed := time.Since(start)
	return elapsed
}

func toBits(seq []byte) []byte {
	for i := 0; i < len(seq); i++ {
		// 'A' => 0, 'C' => 1, 'T' => 2, 'G' => 3
		seq[i] = seq[i] >> 1 & 3
	}
	return seq
}

func readSequence(prefix string, input []byte) (data []byte) {
	in, lineCount := findSequence(prefix, input)
	data = make([]byte, 0, lineCount*61)
	for {
		line, err := in.ReadSlice('\n')
		if len(line) <= 1 || line[0] == '>' {
			break
		}

		last := len(line) - 1
		if line[last] == '\n' {
			line = line[0:last]
		}
		data = append(data, line...)

		if err != nil {
			break
		}
	}
	return
}

func findSequence(prefix string, input []byte) (in *bufio.Reader, lineCount int) {
	pfx := []byte(prefix)
	in = bufio.NewReader(bytes.NewReader(input))
	for {
		line, err := in.ReadSlice('\n')
		if err != nil {
			panic("read error")
		}
		lineCount++
		if line[0] == '>' && bytes.HasPrefix(line, pfx) {
			break
		}
	}
	return
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


func KNucleotideBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(kNucleotideParams, repetitions, TimeKNucleotide)
	base_results := benchmark.TimeImpl(kNucleotideParams, repetitions, TimeKNucleotideBase)
	return scribble_results, base_results
}
