package recTree

import "sync"

type MsgTree interface {
	isMsg_Tree()
}

type Call_M_Tree struct {
	ch_W1_M chan MsgTree
	ch_W2_M chan MsgTree
}

func (lbl Call_M_Tree) isMsg_Tree() {

}

type Next int

func (lbl Next) isMsg_Tree() {

}

type Ctx_Tree_W2 interface {
	Recv_M_Tree_Next_3(x_8 Next)
	Init_W1_Tree_Ctx_2() Ctx_Tree_W1
	Init_W2_Tree_Ctx_2() Ctx_Tree_W2
	Recv_M_Tree_Next_4(x_10 Next)
	Send_W1_Tree_Next_3() Next
	Send_W2_Tree_Next_3() Next
	End()
}

type Ctx_Tree_W1 interface {
	Recv_M_Tree_Next(x_3 Next)
	Init_W1_Tree_Ctx() Ctx_Tree_W1
	Init_W2_Tree_Ctx() Ctx_Tree_W2
	Recv_M_Tree_Next_2(x_5 Next)
	Send_W1_Tree_Next_2() Next
	Send_W2_Tree_Next_2() Next
	End()
}

type Ctx_Tree_M interface {
	Send_W1_Tree_Next() Next
	Send_W2_Tree_Next() Next
	End()
}

func Tree_W2(ctx_5 Ctx_Tree_W2, wg *sync.WaitGroup, ch_W2_M_3, ch_W2_W2_2 chan MsgTree) {
	defer wg.Done()
	x_8 := (<-ch_W2_M_3).(Next)
	ctx_5.Recv_M_Tree_Next_3(x_8)
	ch_W2_M_4 := make(chan MsgTree, 1)
	ch_W1_M_4 := make(chan MsgTree, 1)
	ch_W2_W2_2 <- Call_M_Tree{ch_W1_M_4, ch_W2_M_4}
	ch_W1_W1_3 := make(chan MsgTree, 1)
	ctx_6 := ctx_5.Init_W1_Tree_Ctx_2()
	wg.Add(1)
	go Tree_W1(ctx_6, wg, ch_W1_M_4, ch_W1_W1_3)
	ch_W2_W2_3 := make(chan MsgTree, 1)
	ctx_7 := ctx_5.Init_W2_Tree_Ctx_2()
	wg.Add(1)
	go Tree_W2(ctx_7, wg, ch_W2_M_4, ch_W2_W2_3)
	x_9 := (<-ch_W2_W2_2).(Call_M_Tree)
X:
	for {
		x_10 := (<-ch_W2_M_3).(Next)
		ctx_5.Recv_M_Tree_Next_4(x_10)
		x_11 := ctx_5.Send_W1_Tree_Next_3()
		(x_9).ch_W1_M <- x_11
		x_12 := ctx_5.Send_W2_Tree_Next_3()
		(x_9).ch_W2_M <- x_12
		continue X
	}
}

func Tree_W1(ctx_2 Ctx_Tree_W1, wg *sync.WaitGroup, ch_W1_M_2, ch_W1_W1 chan MsgTree) {
	defer wg.Done()
	x_3 := (<-ch_W1_M_2).(Next)
	ctx_2.Recv_M_Tree_Next(x_3)
	ch_W2_M_2 := make(chan MsgTree, 1)
	ch_W1_M_3 := make(chan MsgTree, 1)
	ch_W1_W1 <- Call_M_Tree{ch_W1_M_3, ch_W2_M_2}
	ch_W1_W1_2 := make(chan MsgTree, 1)
	ctx_3 := ctx_2.Init_W1_Tree_Ctx()
	wg.Add(1)
	go Tree_W1(ctx_3, wg, ch_W1_M_3, ch_W1_W1_2)
	ch_W2_W2 := make(chan MsgTree, 1)
	ctx_4 := ctx_2.Init_W2_Tree_Ctx()
	wg.Add(1)
	go Tree_W2(ctx_4, wg, ch_W2_M_2, ch_W2_W2)
	x_4 := (<-ch_W1_W1).(Call_M_Tree)
X:
	for {
		x_5 := (<-ch_W1_M_2).(Next)
		ctx_2.Recv_M_Tree_Next_2(x_5)
		x_6 := ctx_2.Send_W1_Tree_Next_2()
		(x_4).ch_W1_M <- x_6
		x_7 := ctx_2.Send_W2_Tree_Next_2()
		(x_4).ch_W2_M <- x_7
		continue X
	}
}

func Tree_M(ctx Ctx_Tree_M, wg *sync.WaitGroup, ch_W1_M, ch_W2_M chan MsgTree) {
X:
	for {
		x := ctx.Send_W1_Tree_Next()
		ch_W1_M <- x
		x_2 := ctx.Send_W2_Tree_Next()
		ch_W2_M <- x_2
		continue X
	}
}

func Start(ictx Ctx_Tree_M, ictx_2 Ctx_Tree_W1, ictx_3 Ctx_Tree_W2) {
	var wg sync.WaitGroup
	ch_W2_M_5 := make(chan MsgTree, 1)
	ch_W1_M_5 := make(chan MsgTree, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		Tree_M(ictx, &wg, ch_W1_M_5, ch_W2_M_5)
	}()
	ch_W1_W1_4 := make(chan MsgTree, 1)
	wg.Add(1)
	go Tree_W1(ictx_2, &wg, ch_W1_M_5, ch_W1_W1_4)
	ch_W2_W2_4 := make(chan MsgTree, 1)
	wg.Add(1)
	go Tree_W2(ictx_3, &wg, ch_W2_M_5, ch_W2_W2_4)
	wg.Wait()
}
