package goscr

//import "fmt"

import proto "./boundedPrimeSieve"

type Ctx struct {
    prime int 
    nums []int
}
func (ctx *Ctx) Choice_S_Sieve_SendNums_() proto.Select_S {
    for j, num := range(ctx.nums) {
        if num % ctx.prime != 0 || ctx.prime == 1 {
            (*ctx).nums=((*ctx).nums)[j+1:]
            return proto.Num(num)
        }
    }
    return proto.End{}
}
func (_ *Ctx) End() { }

/* type Ctx_R []int */

func (ctx *Ctx) Recv_S_Sieve_SendNums_Num(v proto.Num) {
    ctx.nums = append(ctx.nums, int(v))
}
func (_ *Ctx) Recv_S_Sieve_SendNums_End(_ proto.End) { }
/* func (ctx *Ctx) End() {} */

type Ctx_M []int
func (ctx *Ctx_M) Recv_Worker_PrimeSieve_Prime(v proto.Prime) {
    *ctx = append(*ctx,int(v))
}
func (ctx *Ctx_M) Init_M_Sieve_Ctx() proto.Ctx_Sieve_M {
    return ctx
}
func (_ *Ctx_M) End_M_Sieve_Ctx(_ proto.Ctx_Sieve_M) { }
func (_ *Ctx_M) Recv_Worker_PrimeSieve_Finish(_ proto.Finish) { }
func (_ *Ctx_M) End() { }
func (ctx *Ctx_M) Recv_W2_Sieve_Prime(v proto.Prime) {
    *ctx = append(*ctx,int(v))
}
func (_ *Ctx_M) Recv_W2_Sieve_Finish(_ proto.Finish) { }


/* type Ctx_W struct */

func (ctx *Ctx) Choice_Worker_PrimeSieve_() proto.Select_Worker {
    if len(ctx.nums) == 0 {
        return proto.Finish{}
    }
    return proto.Prime(ctx.nums[0])
}
func (ctx *Ctx) Init_W2_Sieve_Ctx() proto.Ctx_Sieve_W2 {
    prime := ctx.nums[0]
    ctx.nums = ctx.nums[1:]
    nms := make([]int, 0, len(ctx.nums))
    x := Ctx{prime, nms}
    return &x
}
func (ctx *Ctx) Init_W1_Sieve_Ctx() proto.Ctx_Sieve_W1{
    return ctx
}
func (_ *Ctx) End_W1_Sieve_Ctx(_ proto.Ctx_Sieve_W1) {}
/* func (_ *Ctx) End() {} */


func (ctx *Ctx) Init_R_Sieve_SendNums_Ctx() proto.Ctx_Sieve_SendNums_R {
    return ctx
}
func (_ *Ctx) End_R_Sieve_SendNums_Ctx(_ proto.Ctx_Sieve_SendNums_R) { }
func (ctx *Ctx) Choice_W2_Sieve_() proto.Select_W2 {
    if len(ctx.nums) > 0 {
        return proto.Prime(ctx.nums[0])
    }
    return proto.Finish{}
}
func (ctx *Ctx) Init_W2_Sieve_Ctx_2() proto.Ctx_Sieve_W2 {
    prime := ctx.nums[0]
    ctx.nums = ctx.nums[1:]
    nms := make([]int, 0, len(ctx.nums))
    x := Ctx{prime, nms}
    return &x
}
func (ctx *Ctx) Init_W1_Sieve_Ctx_2() proto.Ctx_Sieve_W1 {
    return ctx
}
func (_ *Ctx) End_W1_Sieve_Ctx_2(_ proto.Ctx_Sieve_W1) { }

func (ctx *Ctx) Init_S_Sieve_SendNums_Ctx() proto.Ctx_Sieve_SendNums_S {
    return ctx
}
func (_ *Ctx) End_S_Sieve_SendNums_Ctx(_ proto.Ctx_Sieve_SendNums_S) { }

func PrimeSieve(n int) []int {
    ctxM := Ctx_M(make([]int, 0, n))
    nums := make([]int, n-2, n-2)
    for i, _ := range(nums) {
        nums[i] = i+3
    }
    ctxW := Ctx{2, nums}
    proto.Start(&ctxM, &ctxW)
    return ctxM
}

//func main() {
//    fmt.Println(PrimeSieve(100))
//}
