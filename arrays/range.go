package arrays

import (
	"fmt"
)

//range is used to iterate over elements in various data structures

func Range() {
	//range used for array
	names := []string{"Pranil", "Octane", "oct10"}
	for index, value := range names {
		fmt.Printf("index: %d, value: %s\n", index, value)
	}

	//range used for strings
	userName := "OctaNil"
	for index, char := range userName {
		fmt.Printf("index: %d\n char: %c,\n", index, char)
	}

	//range used for maps
	// m:= make(map[string]int)
	// m["first"] = 1

	m1 := map[string]int{"first": 1, "second": 2}
	for key, value := range m1 {
		fmt.Printf("key:%s, value:%d\n", key, value)
	}

}
