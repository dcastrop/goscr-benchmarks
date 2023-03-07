package dns

import "sync"

type MsgRecDNSLookup interface {
	isMsg_RecDNSLookup()
}

type MsgIterDNSLookup interface {
	isMsg_IterDNSLookup()
}

type MsgGeneralDNS interface {
	isMsg_GeneralDNS()
}

type Call_res_IterDNSLookup struct {
	ch_dns_res chan MsgIterDNSLookup
	ch_res_dns chan MsgIterDNSLookup
	ch_res_res chan MsgIterDNSLookup
}

func (lbl Call_res_IterDNSLookup) isMsg_GeneralDNS() {

}

type Call_res_RecDNSLookup struct {
	ch_dns_res chan MsgRecDNSLookup
	ch_res_dns chan MsgRecDNSLookup
}

func (lbl Call_res_RecDNSLookup) isMsg_GeneralDNS() {

}

type Done struct {
}

func (lbl Done) isMsg_GeneralDNS() {

}

type IP string

func (lbl IP) isMsg_GeneralDNS() {

}

type Invite_res_IterDNSLookup struct {
}

func (lbl Invite_res_IterDNSLookup) isMsg_GeneralDNS() {

}

type Invite_res_RecDNSLookup struct {
}

func (lbl Invite_res_RecDNSLookup) isMsg_GeneralDNS() {

}

type Query string

func (lbl Query) isMsg_GeneralDNS() {

}

func (lbl Call_res_IterDNSLookup) isMsg_IterDNSLookup() {

}

type DNSIP string

func (lbl DNSIP) isMsg_IterDNSLookup() {

}

func (lbl IP) isMsg_IterDNSLookup() {

}

type IterReq string

func (lbl IterReq) isMsg_IterDNSLookup() {

}

func (lbl Call_res_IterDNSLookup) isMsg_RecDNSLookup() {

}

func (lbl Call_res_RecDNSLookup) isMsg_RecDNSLookup() {

}

func (lbl IP) isMsg_RecDNSLookup() {

}

func (lbl Invite_res_IterDNSLookup) isMsg_RecDNSLookup() {

}

func (lbl Invite_res_RecDNSLookup) isMsg_RecDNSLookup() {

}

type RecReq string

func (lbl RecReq) isMsg_RecDNSLookup() {

}

type Select_dns_2 interface {
	isSelect_dns_2()
}

func (lbl Invite_res_RecDNSLookup) isSelect_dns_2() {

}

func (lbl Invite_res_IterDNSLookup) isSelect_dns_2() {

}

func (lbl IP) isSelect_dns_2() {

}

type Select_dnsRes interface {
	isSelect_dnsRes()
}

func (lbl Invite_res_RecDNSLookup) isSelect_dnsRes() {

}

func (lbl Invite_res_IterDNSLookup) isSelect_dnsRes() {

}

func (lbl IP) isSelect_dnsRes() {

}

type Select_dns interface {
	isSelect_dns()
}

func (lbl IP) isSelect_dns() {

}

func (lbl DNSIP) isSelect_dns() {

}

type Select_app interface {
	isSelect_app()
}

func (lbl Done) isSelect_app() {

}

func (lbl Query) isSelect_app() {

}

type Ctx_RecDNSLookup_res interface {
	Send_dns_RecDNSLookup_RecReq() RecReq
	Recv_dns_RecDNSLookup_IP(x_21 IP)
	End()
}

type Ctx_RecDNSLookup_dns interface {
	Recv_res_RecDNSLookup_RecReq(x_14 RecReq)
	Choice_dns_RecDNSLookup_() Select_dns_2
	Init_dns_RecDNSLookup_Ctx_2() Ctx_RecDNSLookup_dns
	Init_res_RecDNSLookup_Ctx_2() Ctx_RecDNSLookup_res
	End_res_RecDNSLookup_Ctx_2(ctx_12 Ctx_RecDNSLookup_res)
	Send_res_RecDNSLookup_IP() IP
	Init_dns_IterDNSLookup_Ctx_3() Ctx_IterDNSLookup_dns
	Init_res_IterDNSLookup_Ctx_2() Ctx_IterDNSLookup_res
	End_res_IterDNSLookup_Ctx_2(ctx_14 Ctx_IterDNSLookup_res)
	Send_res_RecDNSLookup_IP_2() IP
	End()
}

type Ctx_IterDNSLookup_res interface {
	Send_dns_IterDNSLookup_IterReq() IterReq
	Recv_dns_IterDNSLookup_DNSIP(v_5 DNSIP)
	Init_dns_IterDNSLookup_Ctx_2() Ctx_IterDNSLookup_dns
	Recv_dns_IterDNSLookup_IP(v_5 IP)
	End()
}

type Ctx_IterDNSLookup_dns interface {
	Recv_res_IterDNSLookup_IterReq(x_9 IterReq)
	Choice_dns_IterDNSLookup_() Select_dns
	End()
}

type Ctx_GeneralDNS_dnsRes interface {
	Recv_app_GeneralDNS_Query(v_2 Query)
	Choice_dnsRes_GeneralDNS_() Select_dnsRes
	Init_dns_RecDNSLookup_Ctx() Ctx_RecDNSLookup_dns
	Init_res_RecDNSLookup_Ctx() Ctx_RecDNSLookup_res
	End_res_RecDNSLookup_Ctx(ctx_4 Ctx_RecDNSLookup_res)
	Send_app_GeneralDNS_IP() IP
	Init_dns_IterDNSLookup_Ctx() Ctx_IterDNSLookup_dns
	Init_res_IterDNSLookup_Ctx() Ctx_IterDNSLookup_res
	End_res_IterDNSLookup_Ctx(ctx_6 Ctx_IterDNSLookup_res)
	Send_app_GeneralDNS_IP_2() IP
	Recv_app_GeneralDNS_Done(v_2 Done)
	End()
}

type Ctx_GeneralDNS_app interface {
	Choice_app_GeneralDNS_() Select_app
	Recv_dnsRes_GeneralDNS_IP(x_2 IP)
	End()
}

func RecDNSLookup_res(ctx_15 Ctx_RecDNSLookup_res, wg *sync.WaitGroup, ch_dns_res_9, ch_res_dns_9 chan MsgRecDNSLookup) {
	x_20 := ctx_15.Send_dns_RecDNSLookup_RecReq()
	ch_dns_res_9 <- x_20
	x_21 := (<-ch_res_dns_9).(IP)
	ctx_15.Recv_dns_RecDNSLookup_IP(x_21)
	ctx_15.End()
	return
}

func RecDNSLookup_dns(ctx_10 Ctx_RecDNSLookup_dns, wg *sync.WaitGroup, ch_dns_dns_2, ch_dns_res_6, ch_res_dns_6 chan MsgRecDNSLookup) {
	defer wg.Done()
	x_14 := (<-ch_dns_res_6).(RecReq)
	ctx_10.Recv_res_RecDNSLookup_RecReq(x_14)
	x_15 := ctx_10.Choice_dns_RecDNSLookup_()
	switch v_6 := x_15.(type) {
	case Invite_res_RecDNSLookup:
		ch_res_dns_7 := make(chan MsgRecDNSLookup, 1)
		ch_dns_res_7 := make(chan MsgRecDNSLookup, 1)
		ch_dns_dns_2 <- Call_res_RecDNSLookup{ch_dns_res_7, ch_res_dns_7}
		ch_dns_dns_3 := make(chan MsgRecDNSLookup, 1)
		ctx_11 := ctx_10.Init_dns_RecDNSLookup_Ctx_2()
		wg.Add(1)
		go RecDNSLookup_dns(ctx_11, wg, ch_dns_dns_3, ch_dns_res_7, ch_res_dns_7)
		x_16 := (<-ch_dns_dns_2).(Call_res_RecDNSLookup)
		ctx_12 := ctx_10.Init_res_RecDNSLookup_Ctx_2()
		RecDNSLookup_res(ctx_12, wg, (x_16).ch_dns_res, (x_16).ch_res_dns)
		ctx_10.End_res_RecDNSLookup_Ctx_2(ctx_12)
		x_17 := ctx_10.Send_res_RecDNSLookup_IP()
		ch_res_dns_6 <- x_17
		ctx_10.End()
		return
	case Invite_res_IterDNSLookup:
		ch_res_res_3 := make(chan MsgIterDNSLookup, 1)
		ch_res_dns_8 := make(chan MsgIterDNSLookup, 1)
		ch_dns_res_8 := make(chan MsgIterDNSLookup, 1)
		ch_dns_dns_2 <- Call_res_IterDNSLookup{ch_dns_res_8, ch_res_dns_8, ch_res_res_3}
		ctx_13 := ctx_10.Init_dns_IterDNSLookup_Ctx_3()
		wg.Add(1)
		go IterDNSLookup_dns(ctx_13, wg, ch_dns_res_8, ch_res_dns_8)
		x_18 := (<-ch_dns_dns_2).(Call_res_IterDNSLookup)
		ctx_14 := ctx_10.Init_res_IterDNSLookup_Ctx_2()
		IterDNSLookup_res(ctx_14, wg, (x_18).ch_dns_res, (x_18).ch_res_dns, (x_18).ch_res_res)
		ctx_10.End_res_IterDNSLookup_Ctx_2(ctx_14)
		x_19 := ctx_10.Send_res_RecDNSLookup_IP_2()
		ch_res_dns_6 <- x_19
		ctx_10.End()
		return
	case IP:
		ch_res_dns_6 <- v_6
		ctx_10.End()
		return
	}
}

func IterDNSLookup_res(ctx_8 Ctx_IterDNSLookup_res, wg *sync.WaitGroup, ch_dns_res_4, ch_res_dns_4, ch_res_res_2 chan MsgIterDNSLookup) {
IterDNSLookup_res:
	x_11 := ctx_8.Send_dns_IterDNSLookup_IterReq()
	ch_dns_res_4 <- x_11
	x_12 := <-ch_res_dns_4
	switch v_5 := x_12.(type) {
	case DNSIP:
		ctx_8.Recv_dns_IterDNSLookup_DNSIP(v_5)
		ch_res_dns_5 := make(chan MsgIterDNSLookup, 1)
		ch_dns_res_5 := make(chan MsgIterDNSLookup, 1)
		ch_res_res_2 <- Call_res_IterDNSLookup{ch_dns_res_5, ch_res_dns_5, ch_res_res_2}
		ctx_9 := ctx_8.Init_dns_IterDNSLookup_Ctx_2()
		wg.Add(1)
		go IterDNSLookup_dns(ctx_9, wg, ch_dns_res_5, ch_res_dns_5)
		x_13 := (<-ch_res_res_2).(Call_res_IterDNSLookup)
		ch_res_res_2 = (x_13).ch_res_res
		ch_res_dns_4 = (x_13).ch_res_dns
		ch_dns_res_4 = (x_13).ch_dns_res
		goto IterDNSLookup_res
	case IP:
		ctx_8.Recv_dns_IterDNSLookup_IP(v_5)
		ctx_8.End()
		return
	}
}

func IterDNSLookup_dns(ctx_7 Ctx_IterDNSLookup_dns, wg *sync.WaitGroup, ch_dns_res_3, ch_res_dns_3 chan MsgIterDNSLookup) {
	defer wg.Done()
	x_9 := (<-ch_dns_res_3).(IterReq)
	ctx_7.Recv_res_IterDNSLookup_IterReq(x_9)
	x_10 := ctx_7.Choice_dns_IterDNSLookup_()
	switch v_4 := x_10.(type) {
	case IP:
		ch_res_dns_3 <- v_4
		ctx_7.End()
		return
	case DNSIP:
		ch_res_dns_3 <- v_4
		ctx_7.End()
		return
	}
}

func GeneralDNS_dnsRes(ctx_2 Ctx_GeneralDNS_dnsRes, wg *sync.WaitGroup, ch_app_dnsRes_2, ch_dnsRes_app_2, ch_dnsRes_dnsRes chan MsgGeneralDNS) {
REC:
	for {
		x_3 := <-ch_dnsRes_app_2
		switch v_2 := x_3.(type) {
		case Query:
			ctx_2.Recv_app_GeneralDNS_Query(v_2)
			x_4 := ctx_2.Choice_dnsRes_GeneralDNS_()
			switch v_3 := x_4.(type) {
			case Invite_res_RecDNSLookup:
				ch_res_dns := make(chan MsgRecDNSLookup, 1)
				ch_dns_res := make(chan MsgRecDNSLookup, 1)
				ch_dnsRes_dnsRes <- Call_res_RecDNSLookup{ch_dns_res, ch_res_dns}
				ch_dns_dns := make(chan MsgRecDNSLookup, 1)
				ctx_3 := ctx_2.Init_dns_RecDNSLookup_Ctx()
				wg.Add(1)
				go RecDNSLookup_dns(ctx_3, wg, ch_dns_dns, ch_dns_res, ch_res_dns)
				x_5 := (<-ch_dnsRes_dnsRes).(Call_res_RecDNSLookup)
				ctx_4 := ctx_2.Init_res_RecDNSLookup_Ctx()
				RecDNSLookup_res(ctx_4, wg, (x_5).ch_dns_res, (x_5).ch_res_dns)
				ctx_2.End_res_RecDNSLookup_Ctx(ctx_4)
				x_6 := ctx_2.Send_app_GeneralDNS_IP()
				ch_app_dnsRes_2 <- x_6
				continue REC
			case Invite_res_IterDNSLookup:
				ch_res_res := make(chan MsgIterDNSLookup, 1)
				ch_res_dns_2 := make(chan MsgIterDNSLookup, 1)
				ch_dns_res_2 := make(chan MsgIterDNSLookup, 1)
				ch_dnsRes_dnsRes <- Call_res_IterDNSLookup{ch_dns_res_2, ch_res_dns_2, ch_res_res}
				ctx_5 := ctx_2.Init_dns_IterDNSLookup_Ctx()
				wg.Add(1)
				go IterDNSLookup_dns(ctx_5, wg, ch_dns_res_2, ch_res_dns_2)
				x_7 := (<-ch_dnsRes_dnsRes).(Call_res_IterDNSLookup)
				ctx_6 := ctx_2.Init_res_IterDNSLookup_Ctx()
				IterDNSLookup_res(ctx_6, wg, (x_7).ch_dns_res, (x_7).ch_res_dns, (x_7).ch_res_res)
				ctx_2.End_res_IterDNSLookup_Ctx(ctx_6)
				x_8 := ctx_2.Send_app_GeneralDNS_IP_2()
				ch_app_dnsRes_2 <- x_8
				continue REC
			case IP:
				ch_app_dnsRes_2 <- v_3
				continue REC
			}
		case Done:
			ctx_2.Recv_app_GeneralDNS_Done(v_2)
			ctx_2.End()
			return
		}
	}
}

func GeneralDNS_app(ctx Ctx_GeneralDNS_app, wg *sync.WaitGroup, ch_app_dnsRes, ch_dnsRes_app chan MsgGeneralDNS) {
REC:
	for {
		x := ctx.Choice_app_GeneralDNS_()
		switch v := x.(type) {
		case Done:
			ch_dnsRes_app <- v
			ctx.End()
			return
		case Query:
			ch_dnsRes_app <- v
			x_2 := (<-ch_app_dnsRes).(IP)
			ctx.Recv_dnsRes_GeneralDNS_IP(x_2)
			continue REC
		}
	}
}

func Start(ictx Ctx_GeneralDNS_app, ictx_2 Ctx_GeneralDNS_dnsRes) {
	var wg sync.WaitGroup
	ch_dnsRes_app_3 := make(chan MsgGeneralDNS, 1)
	ch_app_dnsRes_3 := make(chan MsgGeneralDNS, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		GeneralDNS_app(ictx, &wg, ch_app_dnsRes_3, ch_dnsRes_app_3)
	}()
	ch_dnsRes_dnsRes_2 := make(chan MsgGeneralDNS, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		GeneralDNS_dnsRes(ictx_2, &wg, ch_app_dnsRes_3, ch_dnsRes_app_3, ch_dnsRes_dnsRes_2)
	}()
	wg.Wait()
}
