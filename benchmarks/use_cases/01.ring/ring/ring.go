package ring

import "sync"

type MsgRing interface {
	isMsg_Ring()
}

type MsgMain interface {
	isMsg_Main()
}

type Call_P_Ring chan MsgRing

func (lbl Call_P_Ring) isMsg_Main() {

}

type Call_Q_Ring chan MsgRing

func (lbl Call_Q_Ring) isMsg_Main() {

}

type Last int

func (lbl Last) isMsg_Main() {

}

type Next int

func (lbl Next) isMsg_Main() {

}

func (lbl Call_P_Ring) isMsg_Ring() {

}

func (lbl Call_Q_Ring) isMsg_Ring() {

}

func (lbl Last) isMsg_Ring() {

}

func (lbl Next) isMsg_Ring() {

}

type Ctx_Ring_R interface {
	Recv_Q_Ring_Next(x_15 Next)
	Send_P_Ring_Last() Last
	Init_R_Ring_Ctx_2() Ctx_Ring_R
	Recv_Q_Ring_Next_2(x_18 Next)
	Send_R_Ring_Next_2() Next
	End()
}

type Ctx_Ring_Q interface {
	Send_R_Ring_Next() Next
	End()
}

type Ctx_Ring_P interface {
	Recv_R_Ring_Last(x_12 Last)
	End()
}

type Ctx_Main_Q interface {
	Recv_P_Main_Next(x_7 Next)
	Send_P_Main_Last() Last
	Init_R_Ring_Ctx() Ctx_Ring_R
	Recv_P_Main_Next_2(x_10 Next)
	Send_R_Main_Next() Next
	End()
}

type Ctx_Main_P interface {
	Send_Q_Main_Next() Next
	Recv_Q_Main_Last(x_2 Last)
	Send_Q_Main_Next_2() Next
	Recv_R_Main_Last(x_5 Last)
	End()
}

func Ring_R(ctx_6 Ctx_Ring_R, wg *sync.WaitGroup, ch_P_R_3, ch_R_Q_3, ch_R_R_2 chan MsgRing) {
	defer wg.Done()
	x_15 := (<-ch_R_Q_3).(Next)
	ctx_6.Recv_Q_Ring_Next(x_15)
	x_16 := ctx_6.Send_P_Ring_Last()
	ch_P_R_3 <- x_16
	ch_P_R_4 := make(chan MsgRing, 1)
	ch_P_R_3 <- Call_P_Ring(ch_P_R_4)
	ch_R_Q_4 := make(chan MsgRing, 1)
	ch_R_R_2 <- Call_Q_Ring(ch_R_Q_4)
	ch_R_R_3 := make(chan MsgRing, 1)
	ctx_7 := ctx_6.Init_R_Ring_Ctx_2()
	wg.Add(1)
	go Ring_R(ctx_7, wg, ch_P_R_4, ch_R_Q_4, ch_R_R_3)
	x_17 := (<-ch_R_R_2).(Call_Q_Ring)
X:
	for {
		x_18 := (<-ch_R_Q_3).(Next)
		ctx_6.Recv_Q_Ring_Next_2(x_18)
		x_19 := ctx_6.Send_R_Ring_Next_2()
		x_17 <- x_19
		continue X
	}
}

func Ring_Q(ctx_5 Ctx_Ring_Q, wg *sync.WaitGroup, ch_R_Q_2 chan MsgRing) {
X:
	for {
		x_14 := ctx_5.Send_R_Ring_Next()
		ch_R_Q_2 <- x_14
		continue X
	}
}

func Ring_P(ctx_4 Ctx_Ring_P, wg *sync.WaitGroup, ch_P_R_2 chan MsgRing) {
Ring_P:
	x_12 := (<-ch_P_R_2).(Last)
	ctx_4.Recv_R_Ring_Last(x_12)
	x_13 := (<-ch_P_R_2).(Call_P_Ring)
	ch_P_R_2 = x_13
	goto Ring_P
}

func Main_Q(ctx_2 Ctx_Main_Q, wg *sync.WaitGroup, ch_P_Q_2, ch_Q_P_2, ch_Q_Q chan MsgMain) {
	x_7 := (<-ch_Q_P_2).(Next)
	ctx_2.Recv_P_Main_Next(x_7)
	x_8 := ctx_2.Send_P_Main_Last()
	ch_P_Q_2 <- x_8
	ch_P_R := make(chan MsgRing, 1)
	ch_P_Q_2 <- Call_P_Ring(ch_P_R)
	ch_R_Q := make(chan MsgRing, 1)
	ch_Q_Q <- Call_Q_Ring(ch_R_Q)
	ch_R_R := make(chan MsgRing, 1)
	ctx_3 := ctx_2.Init_R_Ring_Ctx()
	wg.Add(1)
	go Ring_R(ctx_3, wg, ch_P_R, ch_R_Q, ch_R_R)
	x_9 := (<-ch_Q_Q).(Call_Q_Ring)
X:
	for {
		x_10 := (<-ch_Q_P_2).(Next)
		ctx_2.Recv_P_Main_Next_2(x_10)
		x_11 := ctx_2.Send_R_Main_Next()
		x_9 <- x_11
		continue X
	}
}

func Main_P(ctx Ctx_Main_P, wg *sync.WaitGroup, ch_P_Q, ch_Q_P chan MsgMain) {
	x := ctx.Send_Q_Main_Next()
	ch_Q_P <- x
	x_2 := (<-ch_P_Q).(Last)
	ctx.Recv_Q_Main_Last(x_2)
	x_3 := (<-ch_P_Q).(Call_P_Ring)
X:
	for {
		x_4 := ctx.Send_Q_Main_Next_2()
		ch_Q_P <- x_4
		x_5 := (<-x_3).(Last)
		ctx.Recv_R_Main_Last(x_5)
		x_6 := (<-x_3).(Call_P_Ring)
		x_3 = x_6
		continue X
	}
}

func Start(ictx Ctx_Main_P, ictx_2 Ctx_Main_Q) {
	var wg sync.WaitGroup
	ch_Q_P_3 := make(chan MsgMain, 1)
	ch_P_Q_3 := make(chan MsgMain, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		Main_P(ictx, &wg, ch_P_Q_3, ch_Q_P_3)
	}()
	ch_Q_Q_2 := make(chan MsgMain, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		Main_Q(ictx_2, &wg, ch_P_Q_3, ch_Q_P_3, ch_Q_Q_2)
	}()
	wg.Wait()
}
