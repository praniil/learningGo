package golangthings

import (
	"fmt"
	"time"
)

func performTask(item string) {
	for i := 0; i < 5; i++ {
		fmt.Println(item, ":", i)
		time.Sleep(100 * time.Millisecond) //prints task every 100 millisecond
	}
}

func Goroutine() {
	start := time.Now()
	go performTask("hello") //spawns another goroutine performTask() and then sleeps for 500 milliseconds
	go performTask("hey")
	time.Sleep(500 * time.Millisecond)
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println(duration)
}
