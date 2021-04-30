//package main
//import (
//    "./boundedFib"
//    "fmt"
//)
//
//func main(){
//        fmt.Println(Fibonacci(1000000))
//}
package goscr

import (
    "./boundedFib"
)

// Declare role contexts
type Res int

func (r *Res) Recv_F3_BoundedFib_Result(v_3 boundedFib.Result) {
    *r = Res(v_3)
}
func (_ *Res) End() {}


type F1 int
func (c *F1) Send_F3_BoundedFib_Fib1() boundedFib.Fib1 {
    return boundedFib.Fib1(*c)
}
func (_ *F1) End() {}

type F2 int
func (c *F2) Send_F3_BoundedFib_Fib2() boundedFib.Fib2 {
    return boundedFib.Fib2(*c)
}
func (_ *F2) Recv_F3_BoundedFib_End(_ boundedFib.End) {
}
func (c *F2) Init_F1_BoundedFib_Ctx() boundedFib.Ctx_BoundedFib_F1 {
    x := F1(*c)
    return &x
} 
func (_ *F2) End_F1_BoundedFib_Ctx(_ boundedFib.Ctx_BoundedFib_F1) {
}
func (_ *F2) End() {}

type F3 struct {
    nth int
    bound int
    fib int
}
func (ctx *F3) Recv_F1_BoundedFib_Fib1(x_4 boundedFib.Fib1){
    ctx.fib = int(x_4)
}
func (ctx *F3) Recv_F2_BoundedFib_Fib2(x_5 boundedFib.Fib2){
    ctx.fib = ctx.fib + int(x_5)
    ctx.nth++
}
func (ctx *F3) Choice_F3_BoundedFib_() boundedFib.Select_F3 {
    if ctx.nth >= ctx.bound {
        return boundedFib.Result(ctx.fib)
    }
    return boundedFib.Invite_Res_BoundedFib{}
}
func (ctx *F3) Init_F3_BoundedFib_Ctx()  boundedFib.Ctx_BoundedFib_F3 {
    return ctx
}
func (ctx *F3) Init_F2_BoundedFib_Ctx() boundedFib.Ctx_BoundedFib_F2 {
    x := F2(ctx.fib) 
    return &x
}
func (_ *F3) End_F2_BoundedFib_Ctx(_  boundedFib.Ctx_BoundedFib_F2) { }
func (_ *F3) Send_F2_BoundedFib_End() boundedFib.End {
    return struct{}{}
}
func (_ *F3) End() {}

func Fibonacci(n int) int {
    var res Res
    f1 := F1(0)
    f2 := F2(1)
    f3 := F3{1, n, 0}
    boundedFib.Start(&res, &f1, &f2, &f3)
    return int(res)
}
