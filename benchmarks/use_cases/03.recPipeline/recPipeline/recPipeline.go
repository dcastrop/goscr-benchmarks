package recPipeline

import "sync"

type MsgPipe interface {
	isMsg_Pipe()
}

type Call_M_Pipe chan MsgPipe

func (lbl Call_M_Pipe) isMsg_Pipe() {

}

type Next int

func (lbl Next) isMsg_Pipe() {

}

type Ctx_Pipe_W interface {
	Recv_M_Pipe_Next(x_2 Next)
	Init_W_Pipe_Ctx() Ctx_Pipe_W
	Recv_M_Pipe_Next_2(x_4 Next)
	Send_W_Pipe_Next_2() Next
	End()
}

type Ctx_Pipe_M interface {
	Send_W_Pipe_Next() Next
	End()
}

func Pipe_W(ctx_2 Ctx_Pipe_W, wg *sync.WaitGroup, ch_W_M_2, ch_W_W chan MsgPipe) {
	defer wg.Done()
	x_2 := (<-ch_W_M_2).(Next)
	ctx_2.Recv_M_Pipe_Next(x_2)
	ch_W_M_3 := make(chan MsgPipe, 1)
	ch_W_W <- Call_M_Pipe(ch_W_M_3)
	ch_W_W_2 := make(chan MsgPipe, 1)
	ctx_3 := ctx_2.Init_W_Pipe_Ctx()
	wg.Add(1)
	go Pipe_W(ctx_3, wg, ch_W_M_3, ch_W_W_2)
	x_3 := (<-ch_W_W).(Call_M_Pipe)
X:
	for {
		x_4 := (<-ch_W_M_2).(Next)
		ctx_2.Recv_M_Pipe_Next_2(x_4)
		x_5 := ctx_2.Send_W_Pipe_Next_2()
		x_3 <- x_5
		continue X
	}
}

func Pipe_M(ctx Ctx_Pipe_M, wg *sync.WaitGroup, ch_W_M chan MsgPipe) {
X:
	for {
		x := ctx.Send_W_Pipe_Next()
		ch_W_M <- x
		continue X
	}
}

func Start(ictx Ctx_Pipe_M, ictx_2 Ctx_Pipe_W) {
	var wg sync.WaitGroup
	ch_W_M_4 := make(chan MsgPipe, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		Pipe_M(ictx, &wg, ch_W_M_4)
	}()
	ch_W_W_3 := make(chan MsgPipe, 1)
	wg.Add(1)
	go Pipe_W(ictx_2, &wg, ch_W_M_4, ch_W_W_3)
	wg.Wait()
}
