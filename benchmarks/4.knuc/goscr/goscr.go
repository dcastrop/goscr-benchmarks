package goscr

import (
    "./knuc"
	"bufio"
	"bytes"
	"fmt"
	"sort"
)

// package main
// 
// import (
//     "./knuc"
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"os"
// 	"sort"
//     "io/ioutil"
// )
// func readFile(file string) []byte {                                                                                    
//     f, err := os.Open(fmt.Sprintf("./data/%s", file))
//     if err != nil {
//         panic("Can't open input file")
//     } 
//     b, err := ioutil.ReadAll(f)
//     if err != nil {
//         panic("Can't read input file")
//     } 
//     err = f.Close()
//     if err != nil {
//         panic("Can't close input file")
//     } 
//     return b
// }     
// 
// func main(){
//     b := readFile("knucleotide-input5000000.txt")
//     s := seqChars(readSequence(">THREE", b))
//     dna := s.toBits()
//     KNucleotide(dna)
// }

// seqString is a sequence of nucleotides as a string: "ACGT..."
type seqString string

// seqChars is a sequence of nucleotides as chars: 'A', 'C', 'G', 'T'...
type seqChars []byte

// seqBits is a sequence of nucleotides as 2 low bits per byte: 0, 1, 3, 2...
type seqBits []byte

// toBits converts *in-place*
func (seq seqChars) toBits() seqBits {
	for i := 0; i < len(seq); i++ {
		// 'A' => 0, 'C' => 1, 'T' => 2, 'G' => 3
		seq[i] = seq[i] >> 1 & 3
	}
	return seqBits(seq)
}

func (seq seqString) seqBits() seqBits {
	return seqChars(seq).toBits()
}

// seq32 is a short (<= 16) sequence of nucleotides in a compact form
// length is not embedded
type seq32 uint32

// seq64 is a short (17..32) sequence of nucleotides in a compact form
// length is not embedded
type seq64 uint64

// seq32 converts a seqBits to a seq32
func (seq seqBits) seq32() seq32 {
	var num seq32
	for _, char := range seq {
		num = num<<2 | seq32(char)
	}
	return num
}

// seq64 converts a seqBits to a seq64
func (seq seqBits) seq64() seq64 {
	var num seq64
	for _, char := range seq {
		num = num<<2 | seq64(char)
	}
	return num
}

// seqString converts a seq32 to a human readable string
func (num seq32) seqString(length int) seqString {
	sequence := make(seqChars, length)
	for i := 0; i < length; i++ {
		sequence[length-i-1] = "ACTG"[num&3]
		num = num >> 2
	}
	return seqString(sequence)
}

type counter uint32

func (dna seqBits) count32(length int) map[seq32]*counter {
	counts := make(map[seq32]*counter)
	key := dna[0 : length-1].seq32()
	mask := seq32(1)<<uint(2*length) - 1
	for index := length - 1; index < len(dna); index++ {
		key = key<<2&mask | seq32(dna[index])
		pointer := counts[key]
		if pointer == nil {
			n := counter(1)
			counts[key] = &n
		} else {
			*pointer++
		}
	}
	return counts
}

func (dna seqBits) count64(length int) map[seq64]*counter {
	counts := make(map[seq64]*counter)
	key := dna[0 : length-1].seq64()
	mask := seq64(1)<<uint(2*length) - 1
	for index := length - 1; index < len(dna); index++ {
		key = key<<2&mask | seq64(dna[index])
		pointer := counts[key]
		if pointer == nil {
			n := counter(1)
			counts[key] = &n
		} else {
			*pointer++
		}
	}
	return counts
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


type seqCount struct {
	seq   seqString
	count counter
}

type seqCounts []seqCount

func (ss seqCounts) Len() int { return len(ss) }

func (ss seqCounts) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }

// Less order descending by count then seq
func (ss seqCounts) Less(i, j int) bool {
	if ss[i].count == ss[j].count {
		return ss[i].seq > ss[j].seq
	}
	return ss[i].count > ss[j].count
}

func frequencyReport(dna seqBits, length int) string {
	counts := dna.count32(length)
	sortedSeqs := make(seqCounts, 0, len(counts))
	for num, pointer := range counts {
		sortedSeqs = append(
			sortedSeqs,
			seqCount{num.seqString(length), *pointer},
		)
	}
	sort.Sort(sortedSeqs)

	var buf bytes.Buffer
	buf.Grow((8 + length) * len(sortedSeqs))
	var scale float32 = 100.0 / float32(len(dna)-length+1)
	for _, sequence := range sortedSeqs {
		buf.WriteString(fmt.Sprintf(
			"%v %.3f\n", sequence.seq,
			float32(sequence.count)*scale),
		)
	}
	return buf.String()
}

func sequenceReport(dna seqBits, sequence seqString) string {
	var pointer *counter
	seq := sequence.seqBits()
	if len(sequence) <= 16 {
		counts := dna.count32(len(sequence))
		pointer = counts[seq.seq32()]
	} else {
		counts := dna.count64(len(sequence))
		pointer = counts[seq.seq64()]
	}
	var sequenceCount counter
	if pointer != nil {
		sequenceCount = *pointer
	}
	return fmt.Sprintf("%v\t%v", sequenceCount, sequence)
}

type Ctx_W struct {
    result string
    dna seqBits
}

func (ctx *Ctx_W) Recv_M_ScheduleJobs_SequenceJob(v_2 knuc.SequenceJob) {
    (*ctx).result = sequenceReport((*ctx).dna, seqString(v_2))
}
func (ctx *Ctx_W) Send_M_ScheduleJobs_SequenceResult() knuc.SequenceResult {
    return knuc.SequenceResult((*ctx).result)
}
func (ctx *Ctx_W) Recv_M_ScheduleJobs_FrequencyJob(v_2 knuc.FrequencyJob) {
    (*ctx).result = frequencyReport((*ctx).dna, int(v_2))
}
func (ctx *Ctx_W) Send_M_ScheduleJobs_FrequencyResult() knuc.FrequencyResult {
    return knuc.FrequencyResult((*ctx).result)
}
func (_ *Ctx_W) Recv_M_ScheduleJobs_Finish(_ knuc.Finish) { }
func (_ *Ctx_W) End () { }

type Ctx_M struct {
    dna seqBits
    jobn int
}
func (ctx *Ctx_M) Choice_M_ScheduleJobs_() knuc.Select_M {
    if (*ctx).jobn >= len(jobs) {
        return knuc.Finish{}
    }
    n := (*ctx).jobn
    (*ctx).jobn++
    return jobs[n]
    
}
func (ctx *Ctx_M) Init_W_ScheduleJobs_Ctx() knuc.Ctx_ScheduleJobs_W {
    x := Ctx_W{"", (*ctx).dna}
    return &x
}
func (_ *Ctx_M) Recv_W_ScheduleJobs_SequenceResult(x knuc.SequenceResult) {
}
func (ctx *Ctx_M) Init_W_ScheduleJobs_Ctx_2() knuc.Ctx_ScheduleJobs_W {
    x := Ctx_W{"", (*ctx).dna}
    return &x
}

func (_ *Ctx_M) Recv_W_ScheduleJobs_FrequencyResult(x knuc.FrequencyResult) {
}
func (_ *Ctx_M) End() {}

var jobs = [...]knuc.Select_M{
	knuc.FrequencyJob(1),
	knuc.FrequencyJob(2),
	knuc.SequenceJob("GGT"),
	knuc.SequenceJob("GGTA"),
	knuc.SequenceJob("GGTATT"),
	knuc.SequenceJob("GGTATTTTAATT"),
	knuc.SequenceJob("GGTATTTTAATTTATAGT"),
}

func KNucleotide(dna seqBits) {
    x := Ctx_M{dna, 0}
    y := Ctx_W{"", dna}
    knuc.Start(&x, &y)
}
