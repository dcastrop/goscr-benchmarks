package benchmark

import "math"

type RunningStat struct {
	MN    int
	MOldM float64
	MNewM float64
	MOldS float64
	MNewS float64
}

func NewRunningStat() *RunningStat {
	return &RunningStat{}
}

func (r *RunningStat) Push(x float64) {
	r.MN++
	if r.MN == 1 {
		r.MOldM = x
		r.MNewM = x
		r.MOldS = 0.0
	} else {
		r.MNewM = r.MOldM + (x-r.MOldM)/float64(r.MN)
		r.MNewS = r.MOldS + (x-r.MOldM)*(x-r.MNewM)

		r.MOldM = r.MNewM
		r.MOldS = r.MNewS
	}
}

func (r *RunningStat) NumDataValues() int {
	return r.MN
}

func (r *RunningStat) Mean() float64 {
	if r.MN > 0 {
		return r.MNewM
	}
	return 0.0
}

func (r *RunningStat) Variance() float64 {
	if r.MN > 1 {
		return r.MNewS / float64(r.MN-1)
	}
	return 0.0
}

func (r *RunningStat) StandardDeviation() float64 {
	return math.Sqrt(r.Variance())
}
