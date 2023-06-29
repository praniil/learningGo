package arrays

import (
	"fmt"
)

func Slice() {
	var s []string
	s = make([]string, 3, 6) // make(type, initial length, capacity i.e size of the underlying array)
	s[0] = "Pranil"
	s[1] = "Ramesh"
	s[2] = "25"
	fmt.Println("emp: ", s, "len: ", len(s), "capacity: ", cap(s))
	//append and copy
	s = append(s, "ram")
	s = append(s, "shyam")
	s = append(s, "ganesh")

	fmt.Println("emp: ", s, "len: ", len(s), "capacity: ", cap(s)) //append = length is increases

	c := make([]string, len(s), 6)
	copy(c, s)
	fmt.Println("emp: ", c, "len: ", len(c), "capacity: ", cap(c))


}
