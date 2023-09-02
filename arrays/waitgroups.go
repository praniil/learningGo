package arrays
//to wait for multiple goroutines to finish we can use a wait group
import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int){
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Millisecond)
	fmt.Printf("workder %d done \n", id)
}

func WaitGroups(){
	var workr sync.WaitGroup
	for i := 0; i < 10; i++ {
		workr.Add(1)
		i1 := i
		go func(){
			defer workr.Done()
			Worker(i1)
		}()
	}
	workr.Wait()
	fmt.Printf("done \n")
}

