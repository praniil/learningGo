package golangthings

//channel is like a pipe that conect concurrent goroutines. Send values to channel from one go routine and receive those values into another goroutine

import (
	"fmt"
	"sync"
	"time"
)

func Channels() {

	//unbuffered channel := when sender tries to send a value on an unbuffered channel it will block until a receiver is ready to receive the value

	/*messages := make(chan string)
	go func() {
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println(msg)
	// fmt.Println(<-messages)*/

	//buffered channel:= has a specified capacity, allowing it to hold a certain number of elements without blocking the sender.
	//when the channel is full, the sender will block until theres available spacein the buffer

	/*
		ch := make(chan string, 3)
		go func() {
			ch <- "dog"
			ch <- "cat"
			ch <- "horse"
		}()

		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	*/

	//Channel Synchronization
	/*
		done := make(chan bool, 1)
		go worker(done)
		Result := <-done
		fmt.Println(Result)
	*/

	//using Waitgroups for multiple go routines to finish
	start := time.Now()
	var wg sync.WaitGroup

	for i := 1; i <= 100; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}
	wg.Wait()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println(duration)
}

func worker(id int) {
	fmt.Printf("Worker %d starting \n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

/*
func worker(done chan bool) {
		fmt.Println("WOrking")
		//Block until we receive a notification from the worker on the channel
		time.Sleep(time.Second)
		fmt.Println("done")
		done <- true
	}
*/
