package base
import "sync"
// package main
// 
// import "sync"
// import "fmt"
// 
// func main() {
//         fmt.Println(Fibonacci(1000000))
// }

type Select interface {
    isSelect()
}

type End struct {}
func (e End) isSelect() {} 
type Call_F1 chan int
func (e Call_F1) isSelect() {} 

type Select_Res interface {
    isSelect_Res()
}

type Result int
func (e Result) isSelect_Res() {} 
type Call_Res chan Select_Res
func (e Call_Res) isSelect_Res() {} 

var nth *int
var bound *int

func bfibF1(fib int, ch_F3_F1 chan int)  {
	ch_F3_F1 <- fib
}


func bfibRes(ch chan Select_Res)  int {
    c := ch
    for {
        x := <-c
        switch v := x.(type) {
        case Result:
            return int(v)
        case Call_Res:
            c = v
            break
        }
    }
}

func bfibF2(fib int, ch_F3_F2 chan int, ch_F2_F3 chan Select)  {
	ch_F3_F2 <- fib
	x_3 := <-ch_F2_F3
	switch v := x_3.(type) {
	case End:
        break
	case Call_F1:
        bfibF1(fib, v)
        break
	}
}

type Call_F2 struct {
    ch32 chan Select
    ch23 chan int
}

func bfibF3(wg *sync.WaitGroup, ch_F2_F3 chan Select,
        ch_F3_F1, ch_F3_F2 chan int, ch_F3_F3 chan Call_F2, ch_Res_F3 chan Select_Res) { 
	defer wg.Done()
	fib := <-ch_F3_F1
    if *nth >= *bound {
        ch_Res_F3 <- Result(fib)
        ch_F2_F3 <- End{}
        return
    }
    *nth++
	fib += <-ch_F3_F2
	ch_Res_F3_2 := make(chan Select_Res,1)
	ch_Res_F3 <- Call_Res(ch_Res_F3_2)
	ch_F3_F1_3 := make(chan int,1)
	ch_F2_F3 <- Call_F1(ch_F3_F1_3)
    ch_F3_F2_3 := make(chan int,1)
	ch_F2_F3_3 := make(chan Select,1)
	ch_F3_F3 <- Call_F2{ch_F2_F3_3, ch_F3_F2_3}
	ch_F3_F3_2 := make(chan Call_F2,1)
    wg.Add(1)
    go bfibF3(wg, ch_F2_F3_3, ch_F3_F1_3,ch_F3_F2_3,ch_F3_F3_2,ch_Res_F3_2) 
    chs := <-ch_F3_F3
    bfibF2(fib, chs.ch23, chs.ch32)
}


func Fibonacci(ubound int)  int {
	var wg sync.WaitGroup
	ch_Res_F3_4 := make(chan Select_Res,1)
    var res int
	wg.Add(1)
	go func () {
		defer wg.Done()
		res = bfibRes(ch_Res_F3_4)
	}()
	ch_F3_F1_4 := make(chan int,1)
	wg.Add(1)
	go func () {
		defer wg.Done()
		bfibF1(0, ch_F3_F1_4)
	}()
	ch_F3_F2_4 := make(chan int,1)
	ch_F2_F3_4 := make(chan Select,1)
	wg.Add(1)
	go func () {
		defer wg.Done()
		bfibF2(1, ch_F3_F2_4, ch_F2_F3_4)
	}()
	ch_F3_F3_3 := make(chan Call_F2,1)
	wg.Add(1)
    curr:=0
    num := ubound
    bound = &num
    nth = &curr
	go bfibF3(&wg,ch_F2_F3_4,ch_F3_F1_4,ch_F3_F2_4,ch_F3_F3_3,ch_Res_F3_4)
	wg.Wait()
    return res
}

