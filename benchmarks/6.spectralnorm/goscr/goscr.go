package goscr

import (
	sn "./spectralnorm"
	"math"
	"runtime"
)

//package main
//
//import (
//    sn "./spectralnorm"
//	"math"
//	"runtime"
//	"fmt"
//)
//func main(){
//    SpectralNorm(4)
//}

var nCPU = runtime.NumCPU()

type Vec []float64

func (v *Vec) Times(ii, n int, u *Vec) {
	ul := len(*u)
	for i := ii; i < n; i++ {
		var vi float64
		for j := 0; j < ul; j++ {
			vi += (*u)[j] / float64(A(i, j))
		}
		(*v)[i] = vi
	}
}

func (v *Vec) TimesTransp(ii, n int, u *Vec) {
	ul := len(*u)
	for i := ii; i < n; i++ {
		var vi float64
		for j := 0; j < ul; j++ {
			vi += (*u)[j] / float64(A(j, i))
		}
		(*v)[i] = vi
	}
}

// func (v Vec) ATimesTransp(u Vec) {
// 	x := make(Vec, len(u))
// 	c := make(chan int, nCPU)
// 	for i := 0; i < nCPU; i++ {
// 		go x.Times(i*len(v)/nCPU, (i+1)*len(v)/nCPU, u, c)
// 	}
// 	wait(c)
// 	for i := 0; i < nCPU; i++ {
// 		go v.TimesTransp(i*len(v)/nCPU, (i+1)*len(v)/nCPU, x, c)
// 	}
// 	wait(c)
// }

func A(i, j int) int {
	return ((i+j)*(i+j+1)/2 + i + 1)
}

// func SpectralNorm(n int) {
// 	u := make(Vec, n)
// 	for i := range u {
// 		u[i] = 1
// 	}
// 	v := make(Vec, n)
// 	for i := 0; i < 10; i++ {
// 		v.ATimesTransp(u)
// 		u.ATimesTransp(v)
// 	}
// 	var vBv, vv float64
// 	for i, vi := range v {
// 		vBv += u[i] * vi
// 		vv += vi * vi
// 	}
// 	_ = math.Sqrt(vBv / vv)
// 	// TODO: Uncomment
// 	// fmt.Printf("%0.9f\n", math.Sqrt(vBv/vv))
// }

func SpectralNorm(n int) {
	u := make(Vec, n)
	for i := range u {
		u[i] = 1
	}
	v := make(Vec, n)
	x := make(Vec, len(u))
	ctxA := CtxM{0, 0, &u, &v, &x, 0}
	sn.Start(&ctxA, &ctxA)
	// fmt.Printf("%0.9f\n", ctxA.res)
}

type CtxM struct {
	i, j    int
	u, v, x *Vec
	res     float64
}

func (c *CtxM) Choice_M_Times_() sn.Select_M_2 {
	if c.j >= nCPU {
		return sn.Finish{}
	}
	i := c.j
	c.j++
	return sn.TimesTask(i)
}
func (c *CtxM) Init_W_Times_Ctx_3() sn.Ctx_Times_W {
	return c
}
func (c *CtxM) Recv_W_Times_TimesResult(x_13 sn.TimesResult) {}

func (c *CtxM) Choice_M_TimesTransp_() sn.Select_M_3 {
	if c.j >= nCPU {
		return sn.Finish{}
	}
	j := c.j
	c.j++
	return sn.TimesTranspTask(j)
}
func (c *CtxM) Init_W_TimesTransp_Ctx_3() sn.Ctx_TimesTransp_W {
	return c
}
func (c *CtxM) Recv_W_TimesTransp_TimesTranspResult(x_18 sn.TimesTranspResult) {}

func (c *CtxM) Choice_M_SpectralNorm_() sn.Select_M {
	// fmt.Println(*c.u, *c.v)
	if c.i >= 9 {
		var vBv, vv float64
		for i, vi := range *c.v {
			vBv += (*c.u)[i] * vi
			vv += vi * vi
		}
		c.res = math.Sqrt(vBv / vv)
		return sn.Finish{}
	}
	x := make(Vec, len(*c.u))
	c.x = &x
	c.i++
	c.j = 1
	return sn.TimesTask(0)
}

func (c *CtxM) Init_W_Times_Ctx() sn.Ctx_Times_W {
	return c
}
func (c *CtxM) Init_M_Times_Ctx() sn.Ctx_Times_M {
	return c
}
func (c *CtxM) End_M_Times_Ctx(ctx_3 sn.Ctx_Times_M) {
	c.j = 0
}
func (c *CtxM) Recv_W_SpectralNorm_TimesResult(x_3 sn.TimesResult) {}
func (c *CtxM) Init_W_TimesTransp_Ctx() sn.Ctx_TimesTransp_W {
	return c
}
func (c *CtxM) Init_M_TimesTransp_Ctx() sn.Ctx_TimesTransp_M {
	return c
}
func (c *CtxM) End_M_TimesTransp_Ctx(ctx_5 sn.Ctx_TimesTransp_M) {
	c.j = 0
	tmp := c.u
	c.u = c.v
	c.v = tmp
}
func (c *CtxM) Init_W_Times_Ctx_2() sn.Ctx_Times_W {
	return c
}
func (c *CtxM) Init_M_Times_Ctx_2() sn.Ctx_Times_M {
	return c
}
func (c *CtxM) End_M_Times_Ctx_2(ctx_7 sn.Ctx_Times_M) {
	c.j = 0
}
func (c *CtxM) Init_W_TimesTransp_Ctx_2() sn.Ctx_TimesTransp_W {
	return c
}
func (c *CtxM) Init_M_TimesTransp_Ctx_2() sn.Ctx_TimesTransp_M {
	return c
}
func (c *CtxM) End_M_TimesTransp_Ctx_2(ctx_9 sn.Ctx_TimesTransp_M) {
	c.j = 0
	tmp := c.u
	c.u = c.v
	c.v = tmp
}

func (c *CtxM) Recv_M_SpectralNorm_TimesTask(t sn.TimesTask) {
	i := int(t)
	c.x.Times(i*len(*c.v)/nCPU, (i+1)*len(*c.v)/nCPU, c.u)
}
func (c *CtxM) Send_M_SpectralNorm_TimesResult() sn.TimesResult {
	return sn.TimesResult{}
}
func (c *CtxM) Recv_M_SpectralNorm_Finish(v_2 sn.Finish) {}
func (c *CtxM) End()                                     {}

func (c *CtxM) Recv_M_Times_TimesTask(t sn.TimesTask) {
	i := int(t)
	c.x.Times(i*len(*c.v)/nCPU, (i+1)*len(*c.v)/nCPU, c.u)
}
func (c *CtxM) Send_M_Times_TimesResult() sn.TimesResult {
	return sn.TimesResult{}
}
func (c *CtxM) Recv_M_Times_Finish(v_4 sn.Finish) {}

func (c *CtxM) Recv_M_TimesTransp_TimesTranspTask(t sn.TimesTranspTask) {
	i := int(t)
	c.v.TimesTransp(i*len(*c.v)/nCPU, (i+1)*len(*c.v)/nCPU, c.x)
}
func (c *CtxM) Send_M_TimesTransp_TimesTranspResult() sn.TimesTranspResult {
	return sn.TimesTranspResult{}
}
func (c *CtxM) Recv_M_TimesTransp_Finish(v_6 sn.Finish) {}
