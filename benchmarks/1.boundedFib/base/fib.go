package base

func fstSender(fib int, sendChan chan int) {
	sendChan <- fib

}

func sndSender(fib int, sendChan1, sendChan2 chan int, stopChan chan bool) {
	sendChan1 <- fib
	stop := <-stopChan
	if !stop {
		fstSender(fib, sendChan2)
	}
}

func fibCalc(fib1Chan, fib2Chan, sendChan1, sendChan2, resChan chan int,
	prevStopChan, stopChan chan bool, returnFib bool) {
	fib1 := <-fib1Chan
	fib2 := <-fib2Chan
	fib := fib1 + fib2
	prevStopChan <- returnFib // If worker is returning a result, prev worker should stop
	if returnFib {
		resChan <- fib
	} else {
		sndSender(fib, sendChan1, sendChan2, stopChan)
	}
}

func Fibonacci(n int) int {
	if n < 3 {
		panic("n should always be > 2")
	}

	resChan := make(chan int, 1)
	fstChan := make(chan int, 1)
	sndChan1 := make(chan int, 1)
	sndChan2 := make(chan int, 1)
	stopChan := make(chan bool, 1)

	go fstSender(1, fstChan)
	go sndSender(1, sndChan1, sndChan2, stopChan)
	for i := 3; i <= n; i++ {
		nextFibChan1 := make(chan int, 1)
		nextFibChan2 := make(chan int, 1)
		nextStopChan := make(chan bool, 1)
		go fibCalc(fstChan, sndChan1, nextFibChan1, nextFibChan2, resChan, stopChan, nextStopChan, i >= n)
		stopChan = nextStopChan
		fstChan = sndChan2
		sndChan1 = nextFibChan1
		sndChan2 = nextFibChan2
	}
	return <-resChan
}
