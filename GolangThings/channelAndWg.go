package golangthings

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func fibo(n int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	if n <= 1 {
		ch <- n
		return
	}

	ch1 := make(chan int)
	ch2 := make(chan int)
	subWG := sync.WaitGroup{} // WaitGroup for the sub-goroutines

	subWG.Add(2)
	go fibo(n-1, ch1, &subWG)
	go fibo(n-2, ch2, &subWG)

	go func() {
		subWG.Wait()
		close(ch1)
		close(ch2)
	}()

	result1, result2 := <-ch1, <-ch2
	fmt.Println(result1, result2)
	ch <- result1 + result2
}

func Fibonacci() {
	startTime := time.Now()
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	n := 10
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1) //it means that one goroutine has been added to the waitGroup
	go fibo(n, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()
	result := <-ch
	elapsedTime := time.Since(startTime)

	fmt.Printf("Fibonacci(%d) = %d\n", n, result)
	fmt.Printf("Time taken: %s\n", elapsedTime)
	fmt.Printf("Number of CPU cores used: %d\n", numCPU)
}
