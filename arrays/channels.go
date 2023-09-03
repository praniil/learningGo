package arrays

import (
	"fmt"

	
)

func Channels(){
	chanl := make(chan string)
	
	go func() { chanl <- "ping"} ()

	msg := <- chanl
	fmt.Println(msg)
	

}