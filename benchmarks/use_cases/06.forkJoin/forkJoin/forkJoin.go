package forkJoin

import "sync"

type MsgForkJoin interface {
	isMsg_ForkJoin()
}

type Call_S_ForkJoin struct {
	ch_S_S chan MsgForkJoin
	ch_W_S chan MsgForkJoin
}

func (lbl Call_S_ForkJoin) isMsg_ForkJoin() {

}

type SubTask int

func (lbl SubTask) isMsg_ForkJoin() {

}

type Task string

func (lbl Task) isMsg_ForkJoin() {

}

type Select_S interface {
	isSelect_S()
}

func (lbl SubTask) isSelect_S() {

}

func (lbl Task) isSelect_S() {

}

type Ctx_ForkJoin_W interface {
	Recv_S_ForkJoin_Task(v_2 Task)
	Recv_S_ForkJoin_SubTask(v_2 SubTask)
	End()
}

type Ctx_ForkJoin_S interface {
	Choice_S_ForkJoin_() Select_S
	Init_W_ForkJoin_Ctx() Ctx_ForkJoin_W
	End()
}

func ForkJoin_W(ctx_3 Ctx_ForkJoin_W, wg *sync.WaitGroup, ch_W_S_3 chan MsgForkJoin) {
	defer wg.Done()
	x_3 := <-ch_W_S_3
	switch v_2 := x_3.(type) {
	case Task:
		ctx_3.Recv_S_ForkJoin_Task(v_2)
		ctx_3.End()
		return
	case SubTask:
		ctx_3.Recv_S_ForkJoin_SubTask(v_2)
		ctx_3.End()
		return
	}
}

func ForkJoin_S(ctx Ctx_ForkJoin_S, wg *sync.WaitGroup, ch_S_S, ch_W_S chan MsgForkJoin) {
ForkJoin_S:
	x := ctx.Choice_S_ForkJoin_()
	switch v := x.(type) {
	case SubTask:
		ch_W_S <- v
		ch_W_S_2 := make(chan MsgForkJoin, 1)
		ch_S_S <- Call_S_ForkJoin{ch_S_S, ch_W_S_2}
		ctx_2 := ctx.Init_W_ForkJoin_Ctx()
		wg.Add(1)
		go ForkJoin_W(ctx_2, wg, ch_W_S_2)
		x_2 := (<-ch_S_S).(Call_S_ForkJoin)
		ch_W_S = (x_2).ch_W_S
		ch_S_S = (x_2).ch_S_S
		goto ForkJoin_S
	case Task:
		ch_W_S <- v
		ctx.End()
		return
	}
}

func Start(ictx Ctx_ForkJoin_S, ictx_2 Ctx_ForkJoin_W) {
	var wg sync.WaitGroup
	ch_W_S_4 := make(chan MsgForkJoin, 1)
	ch_S_S_2 := make(chan MsgForkJoin, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ForkJoin_S(ictx, &wg, ch_S_S_2, ch_W_S_4)
	}()
	wg.Add(1)
	go ForkJoin_W(ictx_2, &wg, ch_W_S_4)
	wg.Wait()
}
