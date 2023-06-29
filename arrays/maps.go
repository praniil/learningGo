package arrays

import (
	"fmt"
)

// maps : golangs built-in associative data type
func Maps() {
	//we can create empty map using builtin make:
	//make(map[key-type]val-type)
	//name[key] = val syntax
	// m:= make(map[string]int)
	// m["first"] = 5
	// m["second"] = 6
	// fmt.Println(m)

	// var var1 int = m["first"]
	// var var2 int = m["second"]
	// fmt.Println(var1)
	// fmt.Println(var2)

	//another way of defining map

	m1 := map[string]int{"first1": 1, "second2": 2}
	fmt.Println(m1["first1"])

	length := len(m1) //gives the length of map
	fmt.Println(length)

	delete(m1, "first1") //deletes the key/value pair of map
	fmt.Println(m1["first1"])

	m1["first1"] = 20 // adding the first key
	fmt.Println(m1["first1"])

	//checking if a key exists in the map

}
