package base

// import "fmt"

func FirstWorker(nChan, filterPrime <-chan int, resChan chan<- int) {
	ubound := <-nChan
	firstPrime := <-filterPrime
	primes := genPrimes(ubound, firstPrime)
	SieveChoice(primes, resChan)
}

func SieveWorker(primeChan, primesIn <-chan int, resChan chan<- int) {
	filterPrime := <-primeChan
	var primes []int
	for prime := range primesIn {
		if prime%filterPrime != 0 {
			primes = append(primes, prime)
		}
	}
	SieveChoice(primes, resChan)
}

func SieveChoice(primes []int, resChan chan<- int) {
	if len(primes) == 0 {
		close(resChan)
	} else {
		resChan <- primes[0]
		filterPrimeChan := make(chan int, 1)
		sendPrimes := make(chan int, 1)
		go SieveWorker(filterPrimeChan, sendPrimes, resChan)
		forwardPrimes(primes, filterPrimeChan, sendPrimes)
	}
}

func genPrimes(ubound int, filterPrime int) []int {
	var primes []int
	for prime := 3; prime <= ubound; prime++ {
		if prime%filterPrime != 0 {
			primes = append(primes, prime)
		}
	}
	return primes
}

func forwardPrimes(primes []int, primeChan, sendPrimes chan<- int) {
	primeChan <- primes[0]
	for i := 1; i < len(primes); i++ {
		sendPrimes <- primes[i]
	}
	close(sendPrimes)
}
func PrimeSieve(n int) []int {
	if n < 2 {
		panic("n should be >= 2")
	}
	sendUbound := make(chan int, 1)
	recvPrimes := make(chan int, 1)
	filterPrime := make(chan int, 1)

	go FirstWorker(sendUbound, filterPrime, recvPrimes)
	sendUbound <- n
	filterPrime <- 2
	primes := []int{2}
	for prime := range recvPrimes {
		primes = append(primes, prime)
	}
	return primes
}
//func main() {                                                                                                          
//    fmt.Println(PrimeSieve(100))
//}


