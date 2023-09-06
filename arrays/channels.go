package arrays

import (
	"fmt"
	"sync"
)

// func Channels(){
// 	chanl := make(chan string)

// 	go func() { chanl <- "ping"} ()

// 	msg := <- chanl
// 	fmt.Println(msg)

// }

func Channels(){
	ch := make (chan int)

	var wg sync.WaitGroup

	for i:= 0; i < 4; i++ {
		wg.Add(1)
		go func (id int){
			defer wg.Done()
			result := id
			ch <- result
		}(i)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for value := range ch {
		fmt.Println("Received: ", value)
	}
}