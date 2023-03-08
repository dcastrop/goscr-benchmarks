package recForkJoin

import "sync"

type MsgForkJoin interface {
	isMsg_ForkJoin()
}

type Call_S_ForkJoin chan MsgForkJoin

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
	Init_W_ForkJoin_Ctx() Ctx_ForkJoin_W
	Recv_S_ForkJoin_Task_2(v_3 Task)
	Send_W_ForkJoin_Task() Task
	Recv_S_ForkJoin_SubTask_2(v_3 SubTask)
	Send_W_ForkJoin_SubTask() SubTask
	End()
}

type Ctx_ForkJoin_S interface {
	Choice_S_ForkJoin_() Select_S
	End()
}

func ForkJoin_W(ctx_2 Ctx_ForkJoin_W, wg *sync.WaitGroup, ch_W_S_2, ch_W_W chan MsgForkJoin) {
	defer wg.Done()
	x_2 := <-ch_W_S_2
	switch v_2 := x_2.(type) {
	case Task:
		ctx_2.Recv_S_ForkJoin_Task(v_2)
		ctx_2.End()
		return
	case SubTask:
		ctx_2.Recv_S_ForkJoin_SubTask(v_2)
		ch_W_S_3 := make(chan MsgForkJoin, 1)
		ch_W_W <- Call_S_ForkJoin(ch_W_S_3)
		ch_W_W_2 := make(chan MsgForkJoin, 1)
		ctx_3 := ctx_2.Init_W_ForkJoin_Ctx()
		wg.Add(1)
		go ForkJoin_W(ctx_3, wg, ch_W_S_3, ch_W_W_2)
		x_3 := (<-ch_W_W).(Call_S_ForkJoin)
	REPEAT:
		for {
			x_4 := <-ch_W_S_2
			switch v_3 := x_4.(type) {
			case Task:
				ctx_2.Recv_S_ForkJoin_Task_2(v_3)
				x_5 := ctx_2.Send_W_ForkJoin_Task()
				x_3 <- x_5
				ctx_2.End()
				return
			case SubTask:
				ctx_2.Recv_S_ForkJoin_SubTask_2(v_3)
				x_6 := ctx_2.Send_W_ForkJoin_SubTask()
				x_3 <- x_6
				continue REPEAT
			}
		}
	}
}

func ForkJoin_S(ctx Ctx_ForkJoin_S, wg *sync.WaitGroup, ch_W_S chan MsgForkJoin) {
REPEAT:
	for {
		x := ctx.Choice_S_ForkJoin_()
		switch v := x.(type) {
		case SubTask:
			ch_W_S <- v
			continue REPEAT
		case Task:
			ch_W_S <- v
			ctx.End()
			return
		}
	}
}

func Start(ictx Ctx_ForkJoin_S, ictx_2 Ctx_ForkJoin_W) {
	var wg sync.WaitGroup
	ch_W_S_4 := make(chan MsgForkJoin, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ForkJoin_S(ictx, &wg, ch_W_S_4)
	}()
	ch_W_W_3 := make(chan MsgForkJoin, 1)
	wg.Add(1)
	go ForkJoin_W(ictx_2, &wg, ch_W_S_4, ch_W_W_3)
	wg.Wait()
}
