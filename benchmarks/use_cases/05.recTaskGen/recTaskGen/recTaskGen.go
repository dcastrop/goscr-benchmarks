package recTaskGen

import "sync"

type MsgDynTaskGen interface {
	isMsg_DynTaskGen()
}

type Call_S_DynTaskGen struct {
	ch_S_W chan MsgDynTaskGen
	ch_W_S chan MsgDynTaskGen
}

func (lbl Call_S_DynTaskGen) isMsg_DynTaskGen() {

}

type LastReq string

func (lbl LastReq) isMsg_DynTaskGen() {

}

type Req string

func (lbl Req) isMsg_DynTaskGen() {

}

type Resp string

func (lbl Resp) isMsg_DynTaskGen() {

}

type Select_S interface {
	isSelect_S()
}

func (lbl Req) isSelect_S() {

}

func (lbl LastReq) isSelect_S() {

}

type Ctx_DynTaskGen_W interface {
	Recv_S_DynTaskGen_LastReq(v_2 LastReq)
	Recv_S_DynTaskGen_Req(v_2 Req)
	Send_S_DynTaskGen_Resp() Resp
	Init_W_DynTaskGen_Ctx() Ctx_DynTaskGen_W
	Recv_S_DynTaskGen_LastReq_2(v_3 LastReq)
	Send_W_DynTaskGen_LastReq() LastReq
	Recv_S_DynTaskGen_Req_2(v_3 Req)
	Send_S_DynTaskGen_Resp_2() Resp
	Send_W_DynTaskGen_Req() Req
	Recv_W_DynTaskGen_Resp_2(x_10 Resp)
	End()
}

type Ctx_DynTaskGen_S interface {
	Choice_S_DynTaskGen_() Select_S
	Recv_W_DynTaskGen_Resp(x_2 Resp)
	End()
}

func DynTaskGen_W(ctx_2 Ctx_DynTaskGen_W, wg *sync.WaitGroup, ch_S_W_2, ch_W_S_2, ch_W_W chan MsgDynTaskGen) {
	defer wg.Done()
	x_3 := <-ch_W_S_2
	switch v_2 := x_3.(type) {
	case LastReq:
		ctx_2.Recv_S_DynTaskGen_LastReq(v_2)
		ctx_2.End()
		return
	case Req:
		ctx_2.Recv_S_DynTaskGen_Req(v_2)
		x_4 := ctx_2.Send_S_DynTaskGen_Resp()
		ch_S_W_2 <- x_4
		ch_W_S_3 := make(chan MsgDynTaskGen, 1)
		ch_S_W_3 := make(chan MsgDynTaskGen, 1)
		ch_W_W <- Call_S_DynTaskGen{ch_S_W_3, ch_W_S_3}
		ch_W_W_2 := make(chan MsgDynTaskGen, 1)
		ctx_3 := ctx_2.Init_W_DynTaskGen_Ctx()
		wg.Add(1)
		go DynTaskGen_W(ctx_3, wg, ch_S_W_3, ch_W_S_3, ch_W_W_2)
		x_5 := (<-ch_W_W).(Call_S_DynTaskGen)
	LOOP:
		for {
			x_6 := <-ch_W_S_2
			switch v_3 := x_6.(type) {
			case LastReq:
				ctx_2.Recv_S_DynTaskGen_LastReq_2(v_3)
				x_7 := ctx_2.Send_W_DynTaskGen_LastReq()
				(x_5).ch_W_S <- x_7
				ctx_2.End()
				return
			case Req:
				ctx_2.Recv_S_DynTaskGen_Req_2(v_3)
				x_8 := ctx_2.Send_S_DynTaskGen_Resp_2()
				ch_S_W_2 <- x_8
				x_9 := ctx_2.Send_W_DynTaskGen_Req()
				(x_5).ch_W_S <- x_9
				x_10 := (<-(x_5).ch_S_W).(Resp)
				ctx_2.Recv_W_DynTaskGen_Resp_2(x_10)
				continue LOOP
			}
		}
	}
}

func DynTaskGen_S(ctx Ctx_DynTaskGen_S, wg *sync.WaitGroup, ch_S_W, ch_W_S chan MsgDynTaskGen) {
LOOP:
	for {
		x := ctx.Choice_S_DynTaskGen_()
		switch v := x.(type) {
		case Req:
			ch_W_S <- v
			x_2 := (<-ch_S_W).(Resp)
			ctx.Recv_W_DynTaskGen_Resp(x_2)
			continue LOOP
		case LastReq:
			ch_W_S <- v
			ctx.End()
			return
		}
	}
}

func Start(ictx Ctx_DynTaskGen_S, ictx_2 Ctx_DynTaskGen_W) {
	var wg sync.WaitGroup
	ch_W_S_4 := make(chan MsgDynTaskGen, 1)
	ch_S_W_4 := make(chan MsgDynTaskGen, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		DynTaskGen_S(ictx, &wg, ch_S_W_4, ch_W_S_4)
	}()
	ch_W_W_3 := make(chan MsgDynTaskGen, 1)
	wg.Add(1)
	go DynTaskGen_W(ictx_2, &wg, ch_S_W_4, ch_W_S_4, ch_W_W_3)
	wg.Wait()
}
