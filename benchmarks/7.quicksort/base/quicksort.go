package base

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

func quickSortWorker(arr []int, resChan chan []int) {
	if len(arr) < SEQ_THRESHOLD {
		seqQuickSort(arr)
	} else {
		pivot := hoarePartition(arr, 0, len(arr)-1)
		leftRes := make(chan []int, 1)
		rightRes := make(chan []int, 1)
		go quickSortWorker(arr[:pivot+1], leftRes)
		go quickSortWorker(arr[pivot+1:], rightRes)
		<-leftRes
		<-rightRes
	}
	resChan <- arr
}

func QuickSort(arr []int) {
	resultChan := make(chan []int, 1)
	go quickSortWorker(arr, resultChan)
	<-resultChan
	// fmt.Println(result)
}
