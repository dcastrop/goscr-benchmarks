package minmax

import "sync"

type Board = []int

type Player = int

type MsgNoughtsAndCrosses interface {
	isMsg_NoughtsAndCrosses()
}

type MsgMinMaxStrategy interface {
	isMsg_MinMaxStrategy()
}

type Call_Master_MinMaxStrategy struct {
	ch_Master_Master chan MsgMinMaxStrategy
	ch_Master_Worker chan MsgMinMaxStrategy
	ch_Worker_Master chan MsgMinMaxStrategy
}

func (lbl Call_Master_MinMaxStrategy) isMsg_MinMaxStrategy() {

}

type CurrState struct {
	ToMove     Player
	CurrPlayer Player
	Board      Board
}

func (lbl CurrState) isMsg_MinMaxStrategy() {

}

type FinalState struct {
	ToMove     Player
	CurrPlayer Player
	Board      Board
}

func (lbl FinalState) isMsg_MinMaxStrategy() {

}

type Invite_Master_MinMaxStrategy struct {
}

func (lbl Invite_Master_MinMaxStrategy) isMsg_MinMaxStrategy() {

}

type Score int

func (lbl Score) isMsg_MinMaxStrategy() {

}

func (lbl Call_Master_MinMaxStrategy) isMsg_NoughtsAndCrosses() {

}

type Draw int

func (lbl Draw) isMsg_NoughtsAndCrosses() {

}

type Move int

func (lbl Move) isMsg_NoughtsAndCrosses() {

}

type Win int

func (lbl Win) isMsg_NoughtsAndCrosses() {

}

type Select_Worker_2 interface {
	isSelect_Worker_2()
}

func (lbl Invite_Master_MinMaxStrategy) isSelect_Worker_2() {

}

func (lbl Score) isSelect_Worker_2() {

}

type Select_Worker interface {
	isSelect_Worker()
}

func (lbl Invite_Master_MinMaxStrategy) isSelect_Worker() {

}

func (lbl Score) isSelect_Worker() {

}

type Select_P2 interface {
	isSelect_P2()
}

func (lbl Win) isSelect_P2() {

}

func (lbl Draw) isSelect_P2() {

}

func (lbl Move) isSelect_P2() {

}

type Select_P1 interface {
	isSelect_P1()
}

func (lbl Win) isSelect_P1() {

}

func (lbl Draw) isSelect_P1() {

}

func (lbl Move) isSelect_P1() {

}

type Select_Master interface {
	isSelect_Master()
}

func (lbl CurrState) isSelect_Master() {

}

func (lbl FinalState) isSelect_Master() {

}

type Ctx_NoughtsAndCrosses_P2 interface {
	Recv_P1_NoughtsAndCrosses_Win(v_7 Win)
	Recv_P1_NoughtsAndCrosses_Draw(v_7 Draw)
	Recv_P1_NoughtsAndCrosses_Move(v_7 Move)
	Init_Worker_MinMaxStrategy_Ctx_5() Ctx_MinMaxStrategy_Worker
	Init_Master_MinMaxStrategy_Ctx_4() Ctx_MinMaxStrategy_Master
	End_Master_MinMaxStrategy_Ctx_4(ctx_13 Ctx_MinMaxStrategy_Master)
	Choice_P2_NoughtsAndCrosses_() Select_P2
	End()
}

type Ctx_NoughtsAndCrosses_P1 interface {
	Init_Worker_MinMaxStrategy_Ctx_4() Ctx_MinMaxStrategy_Worker
	Init_Master_MinMaxStrategy_Ctx_3() Ctx_MinMaxStrategy_Master
	End_Master_MinMaxStrategy_Ctx_3(ctx_10 Ctx_MinMaxStrategy_Master)
	Choice_P1_NoughtsAndCrosses_() Select_P1
	Recv_P2_NoughtsAndCrosses_Win(v_6 Win)
	Recv_P2_NoughtsAndCrosses_Draw(v_6 Draw)
	Recv_P2_NoughtsAndCrosses_Move(v_6 Move)
	End()
}

type Ctx_MinMaxStrategy_Worker interface {
	Recv_Master_MinMaxStrategy_FinalState(v_2 FinalState)
	Choice_Worker_MinMaxStrategy_() Select_Worker
	Init_Worker_MinMaxStrategy_Ctx_2() Ctx_MinMaxStrategy_Worker
	Init_Master_MinMaxStrategy_Ctx() Ctx_MinMaxStrategy_Master
	End_Master_MinMaxStrategy_Ctx(ctx_5 Ctx_MinMaxStrategy_Master)
	Send_Master_MinMaxStrategy_Score() Score
	Recv_Master_MinMaxStrategy_CurrState(v_2 CurrState)
	Choice_Worker_MinMaxStrategy__2() Select_Worker_2
	Init_Worker_MinMaxStrategy_Ctx_3() Ctx_MinMaxStrategy_Worker
	Init_Master_MinMaxStrategy_Ctx_2() Ctx_MinMaxStrategy_Master
	End_Master_MinMaxStrategy_Ctx_2(ctx_7 Ctx_MinMaxStrategy_Master)
	Send_Master_MinMaxStrategy_Score_2() Score
	End()
}

type Ctx_MinMaxStrategy_Master interface {
	Choice_Master_MinMaxStrategy_() Select_Master
	Init_Worker_MinMaxStrategy_Ctx() Ctx_MinMaxStrategy_Worker
	Recv_Worker_MinMaxStrategy_Score(x_3 Score)
	Recv_Worker_MinMaxStrategy_Score_2(x_4 Score)
	End()
}

func NoughtsAndCrosses_P2(ctx_11 Ctx_NoughtsAndCrosses_P2, wg *sync.WaitGroup, ch_P1_P2_2, ch_P2_P1_2, ch_P2_P2 chan MsgNoughtsAndCrosses) {
P1MOVE:
	for {
		x_15 := <-ch_P2_P1_2
		switch v_7 := x_15.(type) {
		case Win:
			ctx_11.Recv_P1_NoughtsAndCrosses_Win(v_7)
			ctx_11.End()
			return
		case Draw:
			ctx_11.Recv_P1_NoughtsAndCrosses_Draw(v_7)
			ctx_11.End()
			return
		case Move:
			ctx_11.Recv_P1_NoughtsAndCrosses_Move(v_7)
			ch_Worker_Master_7 := make(chan MsgMinMaxStrategy, 1)
			ch_Master_Worker_7 := make(chan MsgMinMaxStrategy, 1)
			ch_Master_Master_5 := make(chan MsgMinMaxStrategy, 1)
			ch_P2_P2 <- Call_Master_MinMaxStrategy{ch_Master_Master_5, ch_Master_Worker_7, ch_Worker_Master_7}
			ch_Worker_Worker_6 := make(chan MsgMinMaxStrategy, 1)
			ctx_12 := ctx_11.Init_Worker_MinMaxStrategy_Ctx_5()
			wg.Add(1)
			go MinMaxStrategy_Worker(ctx_12, wg, ch_Master_Worker_7, ch_Worker_Master_7, ch_Worker_Worker_6)
			x_16 := (<-ch_P2_P2).(Call_Master_MinMaxStrategy)
			ctx_13 := ctx_11.Init_Master_MinMaxStrategy_Ctx_4()
			MinMaxStrategy_Master(ctx_13, wg, (x_16).ch_Master_Master, (x_16).ch_Master_Worker, (x_16).ch_Worker_Master)
			ctx_11.End_Master_MinMaxStrategy_Ctx_4(ctx_13)
			x_17 := ctx_11.Choice_P2_NoughtsAndCrosses_()
			switch v_8 := x_17.(type) {
			case Win:
				ch_P1_P2_2 <- v_8
				ctx_11.End()
				return
			case Draw:
				ch_P1_P2_2 <- v_8
				ctx_11.End()
				return
			case Move:
				ch_P1_P2_2 <- v_8
				continue P1MOVE
			}
		}
	}
}

func NoughtsAndCrosses_P1(ctx_8 Ctx_NoughtsAndCrosses_P1, wg *sync.WaitGroup, ch_P1_P1, ch_P1_P2, ch_P2_P1 chan MsgNoughtsAndCrosses) {
P1MOVE:
	for {
		ch_Worker_Master_6 := make(chan MsgMinMaxStrategy, 1)
		ch_Master_Worker_6 := make(chan MsgMinMaxStrategy, 1)
		ch_Master_Master_4 := make(chan MsgMinMaxStrategy, 1)
		ch_P1_P1 <- Call_Master_MinMaxStrategy{ch_Master_Master_4, ch_Master_Worker_6, ch_Worker_Master_6}
		ch_Worker_Worker_5 := make(chan MsgMinMaxStrategy, 1)
		ctx_9 := ctx_8.Init_Worker_MinMaxStrategy_Ctx_4()
		wg.Add(1)
		go MinMaxStrategy_Worker(ctx_9, wg, ch_Master_Worker_6, ch_Worker_Master_6, ch_Worker_Worker_5)
		x_12 := (<-ch_P1_P1).(Call_Master_MinMaxStrategy)
		ctx_10 := ctx_8.Init_Master_MinMaxStrategy_Ctx_3()
		MinMaxStrategy_Master(ctx_10, wg, (x_12).ch_Master_Master, (x_12).ch_Master_Worker, (x_12).ch_Worker_Master)
		ctx_8.End_Master_MinMaxStrategy_Ctx_3(ctx_10)
		x_13 := ctx_8.Choice_P1_NoughtsAndCrosses_()
		switch v_5 := x_13.(type) {
		case Win:
			ch_P2_P1 <- v_5
			ctx_8.End()
			return
		case Draw:
			ch_P2_P1 <- v_5
			ctx_8.End()
			return
		case Move:
			ch_P2_P1 <- v_5
			x_14 := <-ch_P1_P2
			switch v_6 := x_14.(type) {
			case Win:
				ctx_8.Recv_P2_NoughtsAndCrosses_Win(v_6)
				ctx_8.End()
				return
			case Draw:
				ctx_8.Recv_P2_NoughtsAndCrosses_Draw(v_6)
				ctx_8.End()
				return
			case Move:
				ctx_8.Recv_P2_NoughtsAndCrosses_Move(v_6)
				continue P1MOVE
			}
		}
	}
}

func MinMaxStrategy_Worker(ctx_3 Ctx_MinMaxStrategy_Worker, wg *sync.WaitGroup, ch_Master_Worker_3, ch_Worker_Master_3, ch_Worker_Worker_2 chan MsgMinMaxStrategy) {
	defer wg.Done()
	x_5 := <-ch_Worker_Master_3
	switch v_2 := x_5.(type) {
	case FinalState:
		ctx_3.Recv_Master_MinMaxStrategy_FinalState(v_2)
		x_6 := ctx_3.Choice_Worker_MinMaxStrategy_()
		switch v_3 := x_6.(type) {
		case Invite_Master_MinMaxStrategy:
			ch_Worker_Master_4 := make(chan MsgMinMaxStrategy, 1)
			ch_Master_Worker_4 := make(chan MsgMinMaxStrategy, 1)
			ch_Master_Master_2 := make(chan MsgMinMaxStrategy, 1)
			ch_Worker_Worker_2 <- Call_Master_MinMaxStrategy{ch_Master_Master_2, ch_Master_Worker_4, ch_Worker_Master_4}
			ch_Worker_Worker_3 := make(chan MsgMinMaxStrategy, 1)
			ctx_4 := ctx_3.Init_Worker_MinMaxStrategy_Ctx_2()
			wg.Add(1)
			go MinMaxStrategy_Worker(ctx_4, wg, ch_Master_Worker_4, ch_Worker_Master_4, ch_Worker_Worker_3)
			x_7 := (<-ch_Worker_Worker_2).(Call_Master_MinMaxStrategy)
			ctx_5 := ctx_3.Init_Master_MinMaxStrategy_Ctx()
			MinMaxStrategy_Master(ctx_5, wg, (x_7).ch_Master_Master, (x_7).ch_Master_Worker, (x_7).ch_Worker_Master)
			ctx_3.End_Master_MinMaxStrategy_Ctx(ctx_5)
			x_8 := ctx_3.Send_Master_MinMaxStrategy_Score()
			ch_Master_Worker_3 <- x_8
			ctx_3.End()
			return
		case Score:
			ch_Master_Worker_3 <- v_3
			ctx_3.End()
			return
		}
	case CurrState:
		ctx_3.Recv_Master_MinMaxStrategy_CurrState(v_2)
		x_9 := ctx_3.Choice_Worker_MinMaxStrategy__2()
		switch v_4 := x_9.(type) {
		case Invite_Master_MinMaxStrategy:
			ch_Worker_Master_5 := make(chan MsgMinMaxStrategy, 1)
			ch_Master_Worker_5 := make(chan MsgMinMaxStrategy, 1)
			ch_Master_Master_3 := make(chan MsgMinMaxStrategy, 1)
			ch_Worker_Worker_2 <- Call_Master_MinMaxStrategy{ch_Master_Master_3, ch_Master_Worker_5, ch_Worker_Master_5}
			ch_Worker_Worker_4 := make(chan MsgMinMaxStrategy, 1)
			ctx_6 := ctx_3.Init_Worker_MinMaxStrategy_Ctx_3()
			wg.Add(1)
			go MinMaxStrategy_Worker(ctx_6, wg, ch_Master_Worker_5, ch_Worker_Master_5, ch_Worker_Worker_4)
			x_10 := (<-ch_Worker_Worker_2).(Call_Master_MinMaxStrategy)
			ctx_7 := ctx_3.Init_Master_MinMaxStrategy_Ctx_2()
			MinMaxStrategy_Master(ctx_7, wg, (x_10).ch_Master_Master, (x_10).ch_Master_Worker, (x_10).ch_Worker_Master)
			ctx_3.End_Master_MinMaxStrategy_Ctx_2(ctx_7)
			x_11 := ctx_3.Send_Master_MinMaxStrategy_Score_2()
			ch_Master_Worker_3 <- x_11
			ctx_3.End()
			return
		case Score:
			ch_Master_Worker_3 <- v_4
			ctx_3.End()
			return
		}
	}
}

func MinMaxStrategy_Master(ctx Ctx_MinMaxStrategy_Master, wg *sync.WaitGroup, ch_Master_Master, ch_Master_Worker, ch_Worker_Master chan MsgMinMaxStrategy) {
	x := ctx.Choice_Master_MinMaxStrategy_()
	switch v := x.(type) {
	case CurrState:
		ch_Worker_Master <- v
		ch_Worker_Master_2 := make(chan MsgMinMaxStrategy, 1)
		ch_Master_Worker_2 := make(chan MsgMinMaxStrategy, 1)
		ch_Master_Master <- Call_Master_MinMaxStrategy{ch_Master_Master, ch_Master_Worker_2, ch_Worker_Master_2}
		ch_Worker_Worker := make(chan MsgMinMaxStrategy, 1)
		ctx_2 := ctx.Init_Worker_MinMaxStrategy_Ctx()
		wg.Add(1)
		go MinMaxStrategy_Worker(ctx_2, wg, ch_Master_Worker_2, ch_Worker_Master_2, ch_Worker_Worker)
		x_2 := (<-ch_Master_Master).(Call_Master_MinMaxStrategy)
		MinMaxStrategy_Master(ctx, wg, (x_2).ch_Master_Master, (x_2).ch_Master_Worker, (x_2).ch_Worker_Master)
		x_3 := (<-ch_Master_Worker).(Score)
		ctx.Recv_Worker_MinMaxStrategy_Score(x_3)
		ctx.End()
		return
	case FinalState:
		ch_Worker_Master <- v
		x_4 := (<-ch_Master_Worker).(Score)
		ctx.Recv_Worker_MinMaxStrategy_Score_2(x_4)
		ctx.End()
		return
	}
}

func Start(ictx Ctx_NoughtsAndCrosses_P1, ictx_2 Ctx_NoughtsAndCrosses_P2) {
	var wg sync.WaitGroup
	ch_P2_P1_3 := make(chan MsgNoughtsAndCrosses, 1)
	ch_P1_P2_3 := make(chan MsgNoughtsAndCrosses, 1)
	ch_P1_P1_2 := make(chan MsgNoughtsAndCrosses, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		NoughtsAndCrosses_P1(ictx, &wg, ch_P1_P1_2, ch_P1_P2_3, ch_P2_P1_3)
	}()
	ch_P2_P2_2 := make(chan MsgNoughtsAndCrosses, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		NoughtsAndCrosses_P2(ictx_2, &wg, ch_P1_P2_3, ch_P2_P1_3, ch_P2_P2_2)
	}()
	wg.Wait()
}
