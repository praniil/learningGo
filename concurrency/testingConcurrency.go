package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func sum(array []int, size int) int {
	total := 0
	for _, num := range array {
		total = total + num
	}
	return total
}

func sumConcurrent(array []int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	total := 0
	for _, num := range array {
		total += num
	}
	result <- total
}

func Concurrency() {
	// fmt.Println("Con")
	tInitial := time.Now()
	array := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12}
	size := len(array)
	/* without using go concurrency or goroutines */
	total := sum(array, size)
	fmt.Printf("without using concurrency: %d \n", total)
	tFinal := time.Now()
	totalTime := tFinal.Sub(tInitial)
	fmt.Printf("time without using concurrency: %v\n", totalTime)
	//making a worker pool of 4 to divide the task
	timeInitial := time.Now()
	workers := 4
	chunkSize := size / workers
	resultChan := make(chan int, 1)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		if i == workers-1 {
			end = len(array)
		}
		wg.Add(1)
		go sumConcurrent(array[start:end], resultChan, &wg)
	}
	go func() {
		wg.Wait()
		close(resultChan)
	}()
	totalResultFromAllWorkers := 0
	for res := range resultChan {
		totalResultFromAllWorkers += res
	}
	fmt.Printf("concurrent result: %d", totalResultFromAllWorkers)
	timeFinal := time.Now()
	totalConcurrentTime := timeFinal.Sub(timeInitial)
	fmt.Printf("total time while using concurrency feature: %v\n", totalConcurrentTime)
}
