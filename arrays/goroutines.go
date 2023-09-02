package arrays

import (
	"fmt"
	"time"
)

// func sayHello(){
// 	for i:= 0; i<5; i++{
// 		fmt.Println("Hello")
// 		time.Sleep(1000 * time.Millisecond)
// 	}
// }

// func Goroutine(){
// 	go sayHello()
// 	time.Sleep(1000 * time.Millisecond)		//sayHello() call vako 1000*millisecond paxi fmt.Println("Go routine") will be printed
// 	fmt.Println("Go routine")		//hello and Go routine will be printed concurrently
// }

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Goroutine() {
	f("direct")
	go f("goroutine")
	//after Goroutine function is exectuted , done is printed after 1 microsecond and terminate the program, 1 microsecond ma done vanda agadi jj xa tyo execute hunxa
	go f("hellof")

	go func(msg string) {
		fmt.Println(msg)
	}("hello")

	time.Sleep(1 * time.Millisecond)

	fmt.Println("done")
}
