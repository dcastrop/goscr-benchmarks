package taskGen

import "sync"

type MsgDynTaskGen interface {
	isMsg_DynTaskGen()
}

type MsgClientServer interface {
	isMsg_ClientServer()
}

type Call_S_DynTaskGen struct {
	ch_S_S chan MsgDynTaskGen
	ch_S_W chan MsgDynTaskGen
	ch_W_S chan MsgDynTaskGen
}

func (lbl Call_S_DynTaskGen) isMsg_ClientServer() {

}

type Req string

func (lbl Req) isMsg_ClientServer() {

}

type Resp string

func (lbl Resp) isMsg_ClientServer() {

}

func (lbl Call_S_DynTaskGen) isMsg_DynTaskGen() {

}

type LastReq string

func (lbl LastReq) isMsg_DynTaskGen() {

}

func (lbl Req) isMsg_DynTaskGen() {

}

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
	Send_S_DynTaskGen_Resp() Resp
	Recv_S_DynTaskGen_Req(v_2 Req)
	Send_S_DynTaskGen_Resp_2() Resp
	End()
}

type Ctx_DynTaskGen_S interface {
	Choice_S_DynTaskGen_() Select_S
	Init_W_DynTaskGen_Ctx_2() Ctx_DynTaskGen_W
	Recv_W_DynTaskGen_Resp(x_8 Resp)
	Recv_W_DynTaskGen_Resp_2(x_9 Resp)
	End()
}

type Ctx_ClientServer_Server interface {
	Recv_Client_ClientServer_Req(x_3 Req)
	Init_W_DynTaskGen_Ctx() Ctx_DynTaskGen_W
	Init_S_DynTaskGen_Ctx() Ctx_DynTaskGen_S
	End_S_DynTaskGen_Ctx(ctx_4 Ctx_DynTaskGen_S)
	Send_Client_ClientServer_Resp() Resp
	End()
}

type Ctx_ClientServer_Client interface {
	Send_Server_ClientServer_Req() Req
	Recv_Server_ClientServer_Resp(x_2 Resp)
	End()
}

func DynTaskGen_W(ctx_7 Ctx_DynTaskGen_W, wg *sync.WaitGroup, ch_S_W_4, ch_W_S_4 chan MsgDynTaskGen) {
	defer wg.Done()
	x_10 := <-ch_W_S_4
	switch v_2 := x_10.(type) {
	case LastReq:
		ctx_7.Recv_S_DynTaskGen_LastReq(v_2)
		x_11 := ctx_7.Send_S_DynTaskGen_Resp()
		ch_S_W_4 <- x_11
		ctx_7.End()
		return
	case Req:
		ctx_7.Recv_S_DynTaskGen_Req(v_2)
		x_12 := ctx_7.Send_S_DynTaskGen_Resp_2()
		ch_S_W_4 <- x_12
		ctx_7.End()
		return
	}
}

func DynTaskGen_S(ctx_5 Ctx_DynTaskGen_S, wg *sync.WaitGroup, ch_S_S_2, ch_S_W_2, ch_W_S_2 chan MsgDynTaskGen) {
	x_6 := ctx_5.Choice_S_DynTaskGen_()
	switch v := x_6.(type) {
	case Req:
		ch_W_S_2 <- v
		ch_W_S_3 := make(chan MsgDynTaskGen, 1)
		ch_S_W_3 := make(chan MsgDynTaskGen, 1)
		ch_S_S_2 <- Call_S_DynTaskGen{ch_S_S_2, ch_S_W_3, ch_W_S_3}
		ctx_6 := ctx_5.Init_W_DynTaskGen_Ctx_2()
		wg.Add(1)
		go DynTaskGen_W(ctx_6, wg, ch_S_W_3, ch_W_S_3)
		x_7 := (<-ch_S_S_2).(Call_S_DynTaskGen)
		DynTaskGen_S(ctx_5, wg, (x_7).ch_S_S, (x_7).ch_S_W, (x_7).ch_W_S)
		x_8 := (<-ch_S_W_2).(Resp)
		ctx_5.Recv_W_DynTaskGen_Resp(x_8)
		ctx_5.End()
		return
	case LastReq:
		ch_W_S_2 <- v
		x_9 := (<-ch_S_W_2).(Resp)
		ctx_5.Recv_W_DynTaskGen_Resp_2(x_9)
		ctx_5.End()
		return
	}
}

func ClientServer_Server(ctx_2 Ctx_ClientServer_Server, wg *sync.WaitGroup, ch_Client_Server_2, ch_Server_Client_2, ch_Server_Server chan MsgClientServer) {
REPEAT:
	for {
		x_3 := (<-ch_Server_Client_2).(Req)
		ctx_2.Recv_Client_ClientServer_Req(x_3)
		ch_W_S := make(chan MsgDynTaskGen, 1)
		ch_S_W := make(chan MsgDynTaskGen, 1)
		ch_S_S := make(chan MsgDynTaskGen, 1)
		ch_Server_Server <- Call_S_DynTaskGen{ch_S_S, ch_S_W, ch_W_S}
		ctx_3 := ctx_2.Init_W_DynTaskGen_Ctx()
		wg.Add(1)
		go DynTaskGen_W(ctx_3, wg, ch_S_W, ch_W_S)
		x_4 := (<-ch_Server_Server).(Call_S_DynTaskGen)
		ctx_4 := ctx_2.Init_S_DynTaskGen_Ctx()
		DynTaskGen_S(ctx_4, wg, (x_4).ch_S_S, (x_4).ch_S_W, (x_4).ch_W_S)
		ctx_2.End_S_DynTaskGen_Ctx(ctx_4)
		x_5 := ctx_2.Send_Client_ClientServer_Resp()
		ch_Client_Server_2 <- x_5
		continue REPEAT
	}
}

func ClientServer_Client(ctx Ctx_ClientServer_Client, wg *sync.WaitGroup, ch_Client_Server, ch_Server_Client chan MsgClientServer) {
REPEAT:
	for {
		x := ctx.Send_Server_ClientServer_Req()
		ch_Server_Client <- x
		x_2 := (<-ch_Client_Server).(Resp)
		ctx.Recv_Server_ClientServer_Resp(x_2)
		continue REPEAT
	}
}

func Start(ictx Ctx_ClientServer_Client, ictx_2 Ctx_ClientServer_Server) {
	var wg sync.WaitGroup
	ch_Server_Client_3 := make(chan MsgClientServer, 1)
	ch_Client_Server_3 := make(chan MsgClientServer, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ClientServer_Client(ictx, &wg, ch_Client_Server_3, ch_Server_Client_3)
	}()
	ch_Server_Server_2 := make(chan MsgClientServer, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ClientServer_Server(ictx_2, &wg, ch_Client_Server_3, ch_Server_Client_3, ch_Server_Server_2)
	}()
	wg.Wait()
}
