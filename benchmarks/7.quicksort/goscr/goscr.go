package goscr

import qs "./quicksort"

// package main
//
// import qs "./quicksort"
// import "fmt"
//
// func main(){
//     test := []int{9, 3,1,2,3,1,5,1,4,2,1,3,8,6,4,7,9,5}
//     QuickSort(test)
//     fmt.Println(test)
// }
//
// const (
// SEQ_THRESHOLD = 1024
// SEQ_THRESHOLD = 3500
// )

var SEQ_THRESHOLD = 7500

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func hoarePartition(arr []int, low, high int) int {
	pivot := arr[low]
	i, j := low-1, high+1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}
		swap(arr, i, j)
	}
}

func seqQuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := hoarePartition(arr, 0, len(arr)-1)
	seqQuickSort(arr[0 : pivot+1])
	seqQuickSort(arr[pivot+1:])
}

func QuickSort(arr []int) {
	c1 := Ctx(arr)
	c2 := Ctx(arr)
	qs.Start(&c1, &c2)
}

type Ctx []int

// type Ctx_QuickSort_P interface {
func (c *Ctx) Choice_P_QuickSort_() qs.Select_P {
	if len([]int(*c)) < SEQ_THRESHOLD {
		seqQuickSort([]int(*c))
		return qs.Done{}
	} else {
		pivot := hoarePartition([]int(*c), 0, len([]int(*c))-1)
		right := (*c)[pivot+1:]
		*c = (*c)[:pivot+1]
		return qs.Right(right)
	}
}
func (c *Ctx) Init_R_QuickSort_Ctx() qs.Ctx_QuickSort_R {
	nc := Ctx(*c)
	return &nc
}
func (_ *Ctx) Recv_R_QuickSort_Sorted(_ qs.Sorted) {}
func (_ *Ctx) End()                                {}

// type Ctx_QuickSort_R interface {
func (c *Ctx) Recv_P_QuickSort_Right(v_2 qs.Right) {
	*c = ([]int)(v_2)

}
func (c *Ctx) Init_R_QuickSort_Ctx_2() qs.Ctx_QuickSort_R {
	nc := Ctx(*c)
	return &nc
}
func (c *Ctx) Init_P_QuickSort_Ctx() qs.Ctx_QuickSort_P {
	nc := Ctx(*c)
	return &nc
}
func (c *Ctx) End_P_QuickSort_Ctx(ctx_5 qs.Ctx_QuickSort_P) {
}
func (_ *Ctx) Send_P_QuickSort_Sorted() qs.Sorted {
	return qs.Sorted{}
}
func (_ *Ctx) Recv_P_QuickSort_Done(v_2 qs.Done) {}
