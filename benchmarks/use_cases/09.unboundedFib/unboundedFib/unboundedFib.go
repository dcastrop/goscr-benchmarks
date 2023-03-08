package unboundedFib

import "sync"

type MsgFib interface {
	isMsg_Fib()
}

type Call_F1_Fib chan MsgFib

func (lbl Call_F1_Fib) isMsg_Fib() {

}

type Call_F2_Fib struct {
	ch_F2_F3 chan MsgFib
	ch_F3_F2 chan MsgFib
}

func (lbl Call_F2_Fib) isMsg_Fib() {

}

type Call_Res_Fib chan MsgFib

func (lbl Call_Res_Fib) isMsg_Fib() {

}

type Fib1 int

func (lbl Fib1) isMsg_Fib() {

}

type Fib2 int

func (lbl Fib2) isMsg_Fib() {

}

type NextFib int

func (lbl NextFib) isMsg_Fib() {

}

type Ctx_Fib_Res interface {
	Recv_F3_Fib_NextFib(x_8 NextFib)
	End()
}

type Ctx_Fib_F3 interface {
	Recv_F1_Fib_Fib1(x_4 Fib1)
	Recv_F2_Fib_Fib2(x_5 Fib2)
	Send_Res_Fib_NextFib() NextFib
	Init_F3_Fib_Ctx() Ctx_Fib_F3
	Init_F2_Fib_Ctx() Ctx_Fib_F2
	End_F2_Fib_Ctx(ctx_6 Ctx_Fib_F2)
	End()
}

type Ctx_Fib_F2 interface {
	Send_F3_Fib_Fib2() Fib2
	Init_F1_Fib_Ctx() Ctx_Fib_F1
	End_F1_Fib_Ctx(ctx_3 Ctx_Fib_F1)
	End()
}

type Ctx_Fib_F1 interface {
	Send_F3_Fib_Fib1() Fib1
	End()
}

func Fib_Res(ctx_7 Ctx_Fib_Res, wg *sync.WaitGroup, ch_Res_F3_3 chan MsgFib) {
Fib_Res:
	x_8 := (<-ch_Res_F3_3).(NextFib)
	ctx_7.Recv_F3_Fib_NextFib(x_8)
	x_9 := (<-ch_Res_F3_3).(Call_Res_Fib)
	ch_Res_F3_3 = x_9
	goto Fib_Res
}

func Fib_F3(ctx_4 Ctx_Fib_F3, wg *sync.WaitGroup, ch_F2_F3_2, ch_F3_F1_2, ch_F3_F2_2, ch_F3_F3, ch_Res_F3 chan MsgFib) {
	defer wg.Done()
	x_4 := (<-ch_F3_F1_2).(Fib1)
	ctx_4.Recv_F1_Fib_Fib1(x_4)
	x_5 := (<-ch_F3_F2_2).(Fib2)
	ctx_4.Recv_F2_Fib_Fib2(x_5)
	x_6 := ctx_4.Send_Res_Fib_NextFib()
	ch_Res_F3 <- x_6
	ch_Res_F3_2 := make(chan MsgFib, 1)
	ch_Res_F3 <- Call_Res_Fib(ch_Res_F3_2)
	ch_F3_F1_3 := make(chan MsgFib, 1)
	ch_F2_F3_2 <- Call_F1_Fib(ch_F3_F1_3)
	ch_F3_F2_3 := make(chan MsgFib, 1)
	ch_F2_F3_3 := make(chan MsgFib, 1)
	ch_F3_F3 <- Call_F2_Fib{ch_F2_F3_3, ch_F3_F2_3}
	ch_F3_F3_2 := make(chan MsgFib, 1)
	ctx_5 := ctx_4.Init_F3_Fib_Ctx()
	wg.Add(1)
	go Fib_F3(ctx_5, wg, ch_F2_F3_3, ch_F3_F1_3, ch_F3_F2_3, ch_F3_F3_2, ch_Res_F3_2)
	x_7 := (<-ch_F3_F3).(Call_F2_Fib)
	ctx_6 := ctx_4.Init_F2_Fib_Ctx()
	Fib_F2(ctx_6, wg, (x_7).ch_F2_F3, (x_7).ch_F3_F2)
	ctx_4.End_F2_Fib_Ctx(ctx_6)
	ctx_4.End()
	return
}

func Fib_F2(ctx_2 Ctx_Fib_F2, wg *sync.WaitGroup, ch_F2_F3, ch_F3_F2 chan MsgFib) {
	x_2 := ctx_2.Send_F3_Fib_Fib2()
	ch_F3_F2 <- x_2
	x_3 := (<-ch_F2_F3).(Call_F1_Fib)
	ctx_3 := ctx_2.Init_F1_Fib_Ctx()
	Fib_F1(ctx_3, wg, x_3)
	ctx_2.End_F1_Fib_Ctx(ctx_3)
	ctx_2.End()
	return
}

func Fib_F1(ctx Ctx_Fib_F1, wg *sync.WaitGroup, ch_F3_F1 chan MsgFib) {
	x := ctx.Send_F3_Fib_Fib1()
	ch_F3_F1 <- x
	ctx.End()
	return
}

func Start(ictx Ctx_Fib_Res, ictx_2 Ctx_Fib_F1, ictx_3 Ctx_Fib_F2, ictx_4 Ctx_Fib_F3) {
	var wg sync.WaitGroup
	ch_Res_F3_4 := make(chan MsgFib, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		Fib_Res(ictx, &wg, ch_Res_F3_4)
	}()
	ch_F3_F1_4 := make(chan MsgFib, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		Fib_F1(ictx_2, &wg, ch_F3_F1_4)
	}()
	ch_F3_F2_4 := make(chan MsgFib, 1)
	ch_F2_F3_4 := make(chan MsgFib, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		Fib_F2(ictx_3, &wg, ch_F2_F3_4, ch_F3_F2_4)
	}()
	ch_F3_F3_3 := make(chan MsgFib, 1)
	wg.Add(1)
	go Fib_F3(ictx_4, &wg, ch_F2_F3_4, ch_F3_F1_4, ch_F3_F2_4, ch_F3_F3_3, ch_Res_F3_4)
	wg.Wait()
}
