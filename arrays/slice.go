package arrays

import (
	"fmt"
)

func Slice(){
	var s []string
	s = make([]string, 3, 6) // make(type, initial length, capacity i.e size of the underlying array)
	s[0] = "Pranil"
	s[1] = "Ramesh"
	s[2] = "25"
	fmt.Println("emp: ",s , "len: ", len(s), "capacity: ", cap(s))
}