package min-max

import "sync"


type Board = []int

type Player = int

type MsgStandardStrategy interface {
	isMsg_StandardStrategy()
}

type MsgNoughtsAndCrosses interface {
	isMsg_NoughtsAndCrosses()
}

type MsgMinMaxStrategy_EvalBoard interface {
	isMsg_MinMaxStrategy_EvalBoard()
}

type MsgMinMaxStrategy interface {
	isMsg_MinMaxStrategy()
}

type MsgCalcMove interface {
	isMsg_CalcMove()
}

type Call_Master_MinMaxStrategy struct {
	ch_Master_Master chan MsgMinMaxStrategy
	ch_Master_Worker chan MsgMinMaxStrategy
	ch_Worker_Master chan MsgMinMaxStrategy
}

func (lbl Call_Master_MinMaxStrategy) isMsg_CalcMove()  {

}

type Call_P_StandardStrategy struct {

}

func (lbl Call_P_StandardStrategy) isMsg_CalcMove()  {

}

type Invite_Master_MinMaxStrategy struct {

}

func (lbl Invite_Master_MinMaxStrategy) isMsg_CalcMove()  {

}

type Invite_P_StandardStrategy struct {

}

func (lbl Invite_P_StandardStrategy) isMsg_CalcMove()  {

}

func (lbl Call_Master_MinMaxStrategy) isMsg_MinMaxStrategy()  {

}

type Call_W_MinMaxStrategy_EvalBoard struct {

}

func (lbl Call_W_MinMaxStrategy_EvalBoard) isMsg_MinMaxStrategy()  {

}

type CurrState struct {
	ToMove Player
	CurrPlayer Player
	Board Board
}

func (lbl CurrState) isMsg_MinMaxStrategy()  {

}

type FinalState struct {
	ToMove Player
	CurrPlayer Player
	Board Board
}

func (lbl FinalState) isMsg_MinMaxStrategy()  {

}

func (lbl Invite_Master_MinMaxStrategy) isMsg_MinMaxStrategy()  {

}

type Invite_W_MinMaxStrategy_EvalBoard struct {

}

func (lbl Invite_W_MinMaxStrategy_EvalBoard) isMsg_MinMaxStrategy()  {

}

type Score int

func (lbl Score) isMsg_MinMaxStrategy()  {

}

type Call_P_CalcMove chan MsgCalcMove

func (lbl Call_P_CalcMove) isMsg_NoughtsAndCrosses()  {

}

type Draw int

func (lbl Draw) isMsg_NoughtsAndCrosses()  {

}

type Move int

func (lbl Move) isMsg_NoughtsAndCrosses()  {

}

type Win int

func (lbl Win) isMsg_NoughtsAndCrosses()  {

}

type Select_Worker_2 interface {
	isSelect_Worker_2()
}

func (lbl Invite_Master_MinMaxStrategy) isSelect_Worker_2()  {

}

func (lbl Invite_W_MinMaxStrategy_EvalBoard) isSelect_Worker_2()  {

}

type Select_Worker interface {
	isSelect_Worker()
}

func (lbl Invite_Master_MinMaxStrategy) isSelect_Worker()  {

}

func (lbl Invite_W_MinMaxStrategy_EvalBoard) isSelect_Worker()  {

}

type Select_P2 interface {
	isSelect_P2()
}

func (lbl Win) isSelect_P2()  {

}

func (lbl Draw) isSelect_P2()  {

}

func (lbl Move) isSelect_P2()  {

}

type Select_P1 interface {
	isSelect_P1()
}

func (lbl Win) isSelect_P1()  {

}

func (lbl Draw) isSelect_P1()  {

}

func (lbl Move) isSelect_P1()  {

}

type Select_P interface {
	isSelect_P()
}

func (lbl Invite_P_StandardStrategy) isSelect_P()  {

}

func (lbl Invite_Master_MinMaxStrategy) isSelect_P()  {

}

type Select_Master interface {
	isSelect_Master()
}

func (lbl CurrState) isSelect_Master()  {

}

func (lbl FinalState) isSelect_Master()  {

}

type Ctx_StandardStrategy_P interface {
	End()
}

type Ctx_NoughtsAndCrosses_P2 interface {
	Recv_P1_NoughtsAndCrosses_Win(v_8 Win)
	Recv_P1_NoughtsAndCrosses_Draw(v_8 Draw)
	Recv_P1_NoughtsAndCrosses_Move(v_8 Move)
	Init_P_CalcMove_Ctx_2() Ctx_CalcMove_P
	End_P_CalcMove_Ctx_2(ctx_18 Ctx_CalcMove_P)
	Choice_P2_NoughtsAndCrosses_() Select_P2
	End()
}

type Ctx_NoughtsAndCrosses_P1 interface {
	Init_P_CalcMove_Ctx() Ctx_CalcMove_P
	End_P_CalcMove_Ctx(ctx_16 Ctx_CalcMove_P)
	Choice_P1_NoughtsAndCrosses_() Select_P1
	Recv_P2_NoughtsAndCrosses_Win(v_7 Win)
	Recv_P2_NoughtsAndCrosses_Draw(v_7 Draw)
	Recv_P2_NoughtsAndCrosses_Move(v_7 Move)
	End()
}

type Ctx_MinMaxStrategy_EvalBoard_W interface {
	End()
}

type Ctx_MinMaxStrategy_Worker interface {
	Recv_Master_MinMaxStrategy_CurrState(v_3 CurrState)
	Choice_Worker_MinMaxStrategy_() Select_Worker
	Init_Worker_MinMaxStrategy_Ctx_3() Ctx_MinMaxStrategy_Worker
	Init_Master_MinMaxStrategy_Ctx_2() Ctx_MinMaxStrategy_Master
	End_Master_MinMaxStrategy_Ctx_2(ctx_9 Ctx_MinMaxStrategy_Master)
	Send_Master_MinMaxStrategy_Score() Score
	Init_W_MinMaxStrategy_EvalBoard_Ctx() Ctx_MinMaxStrategy_EvalBoard_W
	End_W_MinMaxStrategy_EvalBoard_Ctx(ctx_10 Ctx_MinMaxStrategy_EvalBoard_W)
	Send_Master_MinMaxStrategy_Score_2() Score
	Recv_Master_MinMaxStrategy_FinalState(v_3 FinalState)
	Choice_Worker_MinMaxStrategy__2() Select_Worker_2
	Init_Worker_MinMaxStrategy_Ctx_4() Ctx_MinMaxStrategy_Worker
	Init_Master_MinMaxStrategy_Ctx_3() Ctx_MinMaxStrategy_Master
	End_Master_MinMaxStrategy_Ctx_3(ctx_12 Ctx_MinMaxStrategy_Master)
	Send_Master_MinMaxStrategy_Score_3() Score
	Init_W_MinMaxStrategy_EvalBoard_Ctx_2() Ctx_MinMaxStrategy_EvalBoard_W
	End_W_MinMaxStrategy_EvalBoard_Ctx_2(ctx_13 Ctx_MinMaxStrategy_EvalBoard_W)
	Send_Master_MinMaxStrategy_Score_4() Score
	End()
}

type Ctx_MinMaxStrategy_Master interface {
	Choice_Master_MinMaxStrategy_() Select_Master
	Init_Worker_MinMaxStrategy_Ctx_2() Ctx_MinMaxStrategy_Worker
	Recv_Worker_MinMaxStrategy_Score(x_6 Score)
	Recv_Worker_MinMaxStrategy_Score_2(x_7 Score)
	End()
}

type Ctx_CalcMove_P interface {
	Choice_P_CalcMove_() Select_P
	Init_P_StandardStrategy_Ctx() Ctx_StandardStrategy_P
	End_P_StandardStrategy_Ctx(ctx_2 Ctx_StandardStrategy_P)
	Init_Worker_MinMaxStrategy_Ctx() Ctx_MinMaxStrategy_Worker
	Init_Master_MinMaxStrategy_Ctx() Ctx_MinMaxStrategy_Master
	End_Master_MinMaxStrategy_Ctx(ctx_4 Ctx_MinMaxStrategy_Master)
	End()
}

func StandardStrategy_P(ctx_19 Ctx_StandardStrategy_P, wg *sync.WaitGroup)  {
	ctx_19.End()
	return
}

func NoughtsAndCrosses_P2(ctx_17 Ctx_NoughtsAndCrosses_P2, wg *sync.WaitGroup, ch_P1_P2_2, ch_P2_P1_2, ch_P2_P2 chan MsgNoughtsAndCrosses)  {
	P1MOVE:
	for {
		x_22 := <- ch_P2_P1_2
		switch v_8 := x_22.(type) {
		case Win:
			ctx_17.Recv_P1_NoughtsAndCrosses_Win(v_8)
			ctx_17.End()
			return
		case Draw:
			ctx_17.Recv_P1_NoughtsAndCrosses_Draw(v_8)
			ctx_17.End()
			return
		case Move:
			ctx_17.Recv_P1_NoughtsAndCrosses_Move(v_8)
			ch_P_P_3 := make(chan MsgCalcMove,1)
			ch_P2_P2 <- Call_P_CalcMove(ch_P_P_3)
			x_23 := (<- ch_P2_P2).(Call_P_CalcMove)
			ctx_18 := ctx_17.Init_P_CalcMove_Ctx_2()
			CalcMove_P(ctx_18,wg,x_23)
			ctx_17.End_P_CalcMove_Ctx_2(ctx_18)
			x_24 := ctx_17.Choice_P2_NoughtsAndCrosses_()
			switch v_9 := x_24.(type) {
			case Win:
				ch_P1_P2_2 <- v_9
				ctx_17.End()
				return
			case Draw:
				ch_P1_P2_2 <- v_9
				ctx_17.End()
				return
			case Move:
				ch_P1_P2_2 <- v_9
				continue P1MOVE
			}
		}
	}
}

func NoughtsAndCrosses_P1(ctx_15 Ctx_NoughtsAndCrosses_P1, wg *sync.WaitGroup, ch_P1_P1, ch_P1_P2, ch_P2_P1 chan MsgNoughtsAndCrosses)  {
	P1MOVE:
	for {
		ch_P_P_2 := make(chan MsgCalcMove,1)
		ch_P1_P1 <- Call_P_CalcMove(ch_P_P_2)
		x_19 := (<- ch_P1_P1).(Call_P_CalcMove)
		ctx_16 := ctx_15.Init_P_CalcMove_Ctx()
		CalcMove_P(ctx_16,wg,x_19)
		ctx_15.End_P_CalcMove_Ctx(ctx_16)
		x_20 := ctx_15.Choice_P1_NoughtsAndCrosses_()
		switch v_6 := x_20.(type) {
		case Win:
			ch_P2_P1 <- v_6
			ctx_15.End()
			return
		case Draw:
			ch_P2_P1 <- v_6
			ctx_15.End()
			return
		case Move:
			ch_P2_P1 <- v_6
			x_21 := <- ch_P1_P2
			switch v_7 := x_21.(type) {
			case Win:
				ctx_15.Recv_P2_NoughtsAndCrosses_Win(v_7)
				ctx_15.End()
				return
			case Draw:
				ctx_15.Recv_P2_NoughtsAndCrosses_Draw(v_7)
				ctx_15.End()
				return
			case Move:
				ctx_15.Recv_P2_NoughtsAndCrosses_Move(v_7)
				continue P1MOVE
			}
		}
	}
}

func MinMaxStrategy_EvalBoard_W(ctx_14 Ctx_MinMaxStrategy_EvalBoard_W, wg *sync.WaitGroup)  {
	ctx_14.End()
	return
}

func MinMaxStrategy_Worker(ctx_7 Ctx_MinMaxStrategy_Worker, wg *sync.WaitGroup, ch_Master_Worker_4, ch_Worker_Master_4, ch_Worker_Worker_3 chan MsgMinMaxStrategy)  {
	defer wg.Done()
	x_8 := <- ch_Worker_Master_4
	switch v_3 := x_8.(type) {
	case CurrState:
		ctx_7.Recv_Master_MinMaxStrategy_CurrState(v_3)
		x_9 := ctx_7.Choice_Worker_MinMaxStrategy_()
		switch v_4 := x_9.(type) {
		case Invite_Master_MinMaxStrategy:
			ch_Worker_Master_5 := make(chan MsgMinMaxStrategy,1)
			ch_Master_Worker_5 := make(chan MsgMinMaxStrategy,1)
			ch_Master_Master_3 := make(chan MsgMinMaxStrategy,1)
			ch_Worker_Worker_3 <- Call_Master_MinMaxStrategy{ch_Master_Master_3, ch_Master_Worker_5, ch_Worker_Master_5}
			ch_Worker_Worker_4 := make(chan MsgMinMaxStrategy,1)
			ctx_8 := ctx_7.Init_Worker_MinMaxStrategy_Ctx_3()
			wg.Add(1)
			go MinMaxStrategy_Worker(ctx_8,wg,ch_Master_Worker_5,ch_Worker_Master_5,ch_Worker_Worker_4)
			x_10 := (<- ch_Worker_Worker_3).(Call_Master_MinMaxStrategy)
			ctx_9 := ctx_7.Init_Master_MinMaxStrategy_Ctx_2()
			MinMaxStrategy_Master(ctx_9,wg,(x_10).ch_Master_Master,(x_10).ch_Master_Worker,(x_10).ch_Worker_Master)
			ctx_7.End_Master_MinMaxStrategy_Ctx_2(ctx_9)
			x_11 := ctx_7.Send_Master_MinMaxStrategy_Score()
			ch_Master_Worker_4 <- x_11
			ctx_7.End()
			return
		case Invite_W_MinMaxStrategy_EvalBoard:
			ch_Worker_Worker_3 <- Call_W_MinMaxStrategy_EvalBoard{}
			x_12 := (<- ch_Worker_Worker_3).(Call_W_MinMaxStrategy_EvalBoard)
			ctx_10 := ctx_7.Init_W_MinMaxStrategy_EvalBoard_Ctx()
			MinMaxStrategy_EvalBoard_W(ctx_10,wg)
			ctx_7.End_W_MinMaxStrategy_EvalBoard_Ctx(ctx_10)
			x_13 := ctx_7.Send_Master_MinMaxStrategy_Score_2()
			ch_Master_Worker_4 <- x_13
			ctx_7.End()
			return
		}
	case FinalState:
		ctx_7.Recv_Master_MinMaxStrategy_FinalState(v_3)
		x_14 := ctx_7.Choice_Worker_MinMaxStrategy__2()
		switch v_5 := x_14.(type) {
		case Invite_Master_MinMaxStrategy:
			ch_Worker_Master_6 := make(chan MsgMinMaxStrategy,1)
			ch_Master_Worker_6 := make(chan MsgMinMaxStrategy,1)
			ch_Master_Master_4 := make(chan MsgMinMaxStrategy,1)
			ch_Worker_Worker_3 <- Call_Master_MinMaxStrategy{ch_Master_Master_4, ch_Master_Worker_6, ch_Worker_Master_6}
			ch_Worker_Worker_5 := make(chan MsgMinMaxStrategy,1)
			ctx_11 := ctx_7.Init_Worker_MinMaxStrategy_Ctx_4()
			wg.Add(1)
			go MinMaxStrategy_Worker(ctx_11,wg,ch_Master_Worker_6,ch_Worker_Master_6,ch_Worker_Worker_5)
			x_15 := (<- ch_Worker_Worker_3).(Call_Master_MinMaxStrategy)
			ctx_12 := ctx_7.Init_Master_MinMaxStrategy_Ctx_3()
			MinMaxStrategy_Master(ctx_12,wg,(x_15).ch_Master_Master,(x_15).ch_Master_Worker,(x_15).ch_Worker_Master)
			ctx_7.End_Master_MinMaxStrategy_Ctx_3(ctx_12)
			x_16 := ctx_7.Send_Master_MinMaxStrategy_Score_3()
			ch_Master_Worker_4 <- x_16
			ctx_7.End()
			return
		case Invite_W_MinMaxStrategy_EvalBoard:
			ch_Worker_Worker_3 <- Call_W_MinMaxStrategy_EvalBoard{}
			x_17 := (<- ch_Worker_Worker_3).(Call_W_MinMaxStrategy_EvalBoard)
			ctx_13 := ctx_7.Init_W_MinMaxStrategy_EvalBoard_Ctx_2()
			MinMaxStrategy_EvalBoard_W(ctx_13,wg)
			ctx_7.End_W_MinMaxStrategy_EvalBoard_Ctx_2(ctx_13)
			x_18 := ctx_7.Send_Master_MinMaxStrategy_Score_4()
			ch_Master_Worker_4 <- x_18
			ctx_7.End()
			return
		}
	}
}

func MinMaxStrategy_Master(ctx_5 Ctx_MinMaxStrategy_Master, wg *sync.WaitGroup, ch_Master_Master_2, ch_Master_Worker_2, ch_Worker_Master_2 chan MsgMinMaxStrategy)  {
	x_4 := ctx_5.Choice_Master_MinMaxStrategy_()
	switch v_2 := x_4.(type) {
	case CurrState:
		ch_Worker_Master_2 <- v_2
		ch_Worker_Master_3 := make(chan MsgMinMaxStrategy,1)
		ch_Master_Worker_3 := make(chan MsgMinMaxStrategy,1)
		ch_Master_Master_2 <- Call_Master_MinMaxStrategy{ch_Master_Master_2, ch_Master_Worker_3, ch_Worker_Master_3}
		ch_Worker_Worker_2 := make(chan MsgMinMaxStrategy,1)
		ctx_6 := ctx_5.Init_Worker_MinMaxStrategy_Ctx_2()
		wg.Add(1)
		go MinMaxStrategy_Worker(ctx_6,wg,ch_Master_Worker_3,ch_Worker_Master_3,ch_Worker_Worker_2)
		x_5 := (<- ch_Master_Master_2).(Call_Master_MinMaxStrategy)
		MinMaxStrategy_Master(ctx_5,wg,(x_5).ch_Master_Master,(x_5).ch_Master_Worker,(x_5).ch_Worker_Master)
		x_6 := (<- ch_Master_Worker_2).(Score)
		ctx_5.Recv_Worker_MinMaxStrategy_Score(x_6)
		ctx_5.End()
		return
	case FinalState:
		ch_Worker_Master_2 <- v_2
		x_7 := (<- ch_Master_Worker_2).(Score)
		ctx_5.Recv_Worker_MinMaxStrategy_Score_2(x_7)
		ctx_5.End()
		return
	}
}

func CalcMove_P(ctx Ctx_CalcMove_P, wg *sync.WaitGroup, ch_P_P chan MsgCalcMove)  {
	x := ctx.Choice_P_CalcMove_()
	switch v := x.(type) {
	case Invite_P_StandardStrategy:
		ch_P_P <- Call_P_StandardStrategy{}
		x_2 := (<- ch_P_P).(Call_P_StandardStrategy)
		ctx_2 := ctx.Init_P_StandardStrategy_Ctx()
		StandardStrategy_P(ctx_2,wg)
		ctx.End_P_StandardStrategy_Ctx(ctx_2)
		ctx.End()
		return
	case Invite_Master_MinMaxStrategy:
		ch_Worker_Master := make(chan MsgMinMaxStrategy,1)
		ch_Master_Worker := make(chan MsgMinMaxStrategy,1)
		ch_Master_Master := make(chan MsgMinMaxStrategy,1)
		ch_P_P <- Call_Master_MinMaxStrategy{ch_Master_Master, ch_Master_Worker, ch_Worker_Master}
		ch_Worker_Worker := make(chan MsgMinMaxStrategy,1)
		ctx_3 := ctx.Init_Worker_MinMaxStrategy_Ctx()
		wg.Add(1)
		go MinMaxStrategy_Worker(ctx_3,wg,ch_Master_Worker,ch_Worker_Master,ch_Worker_Worker)
		x_3 := (<- ch_P_P).(Call_Master_MinMaxStrategy)
		ctx_4 := ctx.Init_Master_MinMaxStrategy_Ctx()
		MinMaxStrategy_Master(ctx_4,wg,(x_3).ch_Master_Master,(x_3).ch_Master_Worker,(x_3).ch_Worker_Master)
		ctx.End_Master_MinMaxStrategy_Ctx(ctx_4)
		ctx.End()
		return
	}
}

func Start(ictx Ctx_NoughtsAndCrosses_P1, ictx_2 Ctx_NoughtsAndCrosses_P2)  {
	var wg sync.WaitGroup
	ch_P2_P1_3 := make(chan MsgNoughtsAndCrosses,1)
	ch_P1_P2_3 := make(chan MsgNoughtsAndCrosses,1)
	ch_P1_P1_2 := make(chan MsgNoughtsAndCrosses,1)
	wg.Add(1)
	go func () {
		defer wg.Done()
		NoughtsAndCrosses_P1(ictx,&wg,ch_P1_P1_2,ch_P1_P2_3,ch_P2_P1_3)
	}()
	ch_P2_P2_2 := make(chan MsgNoughtsAndCrosses,1)
	wg.Add(1)
	go func () {
		defer wg.Done()
		NoughtsAndCrosses_P2(ictx_2,&wg,ch_P1_P2_3,ch_P2_P1_3,ch_P2_P2_2)
	}()
	wg.Wait()
}
