package arrays

//to wait for multiple goroutines to finish we can use a wait group
import (
	"fmt"
	"sync"
	"time"
)

func Worker(id int) {
	// defer worgrp.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Millisecond) //1 milisecond samma
	fmt.Printf("workder %d done \n", id)
	//defer keyword is used to schedule a function call to be executed at the end of the surrounding functions scope just before it returns
}

func WaitGroups() {
	var workr sync.WaitGroup
	for i := 0; i < 1000; i++ {
		workr.Add(1)
		go func(in int) {
			defer workr.Done()
			Worker(in)
		}(i)

		// go Worker(i, &workr)		//go routine
	}
	workr.Wait()
	fmt.Printf("done \n")
}
