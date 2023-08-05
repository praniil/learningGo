package arrays

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// func add (a int , b int) int {
// 	return a + b
// }

// func addFloat (a float64, b float64) float64 {
// 	return a + b
// } hectic job
type num interface{
	int | int64 | int32 | float64 | float32
}
//not the ideal
func add [T int| float64] (a T, b T) T {
	return a + b 
}
//OR
//not the ideal

func addnums [T num] (a T, b T) T {
	return a + b
}
//OR

func addAgain [T constraints.Ordered] (a T, b T) T{
	return a + b
}

//golang has 

func Generic () {
	// result := add(1, 3)
	// result1 := addFloat(1.0, 3.2) 
	result := add(1, 3)
	resultAdd := addnums(2, 3)
	resultAddAgain := addAgain(2, 3)

	fmt.Println(result)
	fmt.Println(resultAdd)
	fmt.Println(resultAddAgain)
	
}