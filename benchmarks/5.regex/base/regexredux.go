package base

import (
	"github.com/GRbit/go-pcre"
)

type substitution struct {
	pattern     string
	replacement string
}

var variants = []string{
	"agggtaaa|tttaccct",
	"[cgt]gggtaaa|tttaccc[acg]",
	"a[act]ggtaaa|tttacc[agt]t",
	"ag[act]gtaaa|tttac[agt]ct",
	"agg[act]taaa|ttta[agt]cct",
	"aggg[acg]aaa|ttt[cgt]ccct",
	"agggt[cgt]aa|tt[acg]accct",
	"agggta[cgt]a|t[acg]taccct",
	"agggtaa[cgt]|[acg]ttaccct",
}

var substs = []substitution{
	{"tHa[Nt]", "<4>"},
	{"aND|caN|Ha[DS]|WaS", "<3>"},
	{"a[NSt]|BY", "<2>"},
	{"<[^>]*>", "|"},
	{"\\|[^|][^|]*\\|", "-"},
}

func countMatches(pat string, b []byte) int {
	m := pcre.MustCompileJIT(pat, 0, pcre.STUDY_JIT_COMPILE).Matcher(b, 0)
	n := 0

	for f := m.Matches; f; f = m.Match(b, 0) {
		n++

		b = b[m.Index()[1]:]
	}

	return n
}

func RegexRedux(b []byte) {

	// runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO:Uncomment
	// ilen := len(b)

	// Delete the comment lines and newlines
	b = pcre.
		MustCompileJIT("(>[^\n]*)?\n", 0, pcre.STUDY_JIT_COMPILE).
		ReplaceAll(b, []byte{}, 0)
	// TODO:Uncomment
	// clen := len(b)

	mresults := make([]chan int, len(variants))
	for i := 0; i < len(variants); i++ {
		mresults[i] = make(chan int, 1)

		go func(ch chan int, s string) {
			ch <- countMatches(s, b)
		}(mresults[i], variants[i])
	}

	lenresult := make(chan int, 1)

	go func(b []byte) {
		for i := 0; i < len(substs); i++ {
			b = pcre.
				MustCompileJIT(substs[i].pattern, 0, pcre.STUDY_JIT_COMPILE).
				ReplaceAll(b, []byte(substs[i].replacement), 0)
		}
		lenresult <- len(b)
	}(b)

	for i := 0; i < len(variants); i++ {
		// TODO: Uncomment
		<-mresults[i]
		// fmt.Printf("%s %d\n", variants[i], <-mresults[i])
	}

	// TODO: Uncomment
	<-lenresult
	// fmt.Printf("\n%d\n%d\n%d\n", ilen, clen, <-lenresult)
}
