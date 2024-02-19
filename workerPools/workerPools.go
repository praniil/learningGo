package workerpools

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	// time.Sleep(time.Second * 3)
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2		//j * 2 := in actual scenario computing some input and sending result through the channel
	}

}

func WorkerPool() {
	var wg sync.WaitGroup
	wg.Add(3)
	const numJobs = 9
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for w := 1; w <= 3; w++ {
		go worker(&wg, w, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= numJobs; a++ {
		<-results
	}
	wg.Wait()
}
