package goscr

import (
    "./regex"
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

type Ctx_W int
func (ctx *Ctx_W) Recv_M_Regex_Task(v_2 regex.Task) {
    *ctx = Ctx_W(countMatches(v_2.Pattern, v_2.B))
}
func (ctx *Ctx_W) Send_M_Regex_NumMatches() regex.NumMatches {
    return regex.NumMatches(*ctx)
}
func (ctx *Ctx_W) Recv_M_Regex_CalcLength(b regex.CalcLength) {
		for i := 0; i < len(substs); i++ {
			b = pcre.
				MustCompileJIT(substs[i].pattern, 0, pcre.STUDY_JIT_COMPILE).
				ReplaceAll(b, []byte(substs[i].replacement), 0)
		}
        *ctx = Ctx_W(len(b))
}
func (ctx *Ctx_W) Send_M_Regex_Length() regex.Length {
    return regex.Length(*ctx)
}
func (_ *Ctx_W) End() { }


type Ctx_M struct {
    idx int
    b []byte
}
func (ctx *Ctx_M) Choice_M_Regex_() regex.Select_M {
    if (*ctx).idx >= len(variants) {
        return regex.CalcLength((*ctx).b)
    }
    n := (*ctx).idx
    (*ctx).idx++
    return regex.Task{(*ctx).b, variants[n]}
}
func (_ *Ctx_M) Init_W_Regex_Ctx() regex.Ctx_Regex_W {
    n := Ctx_W(0)
    return &n
}
func (_ *Ctx_M) Recv_W_Regex_NumMatches(x_3 regex.NumMatches) {
}
func (_ *Ctx_M) Recv_W_Regex_Length(x_4 regex.Length) {
}
func (_ *Ctx_M) End() {}


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
    ctx_M := Ctx_M{0, b}
    ctx_W := Ctx_W(0)

    regex.Start(&ctx_M, &ctx_W)
}
