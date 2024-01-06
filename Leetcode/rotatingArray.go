package leetcode

import (
	"fmt"
)

func exchange(array1 [10]int, number int) [10]int {
	var index int = 0
	fmt.Println(array1)
	for number < len(array1) {
		array1[index] = array1[number]
		number++
		// fmt.Println(number)
		index++
		// fmt.Println(index)
		// fmt.Println(array1)
	}

	return array1
}

func RoatatingArray() {
	array1 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array1)
	fmt.Printf("enter the number: ")
	var number int
	fmt.Scanf("%d", &number)
	var number2 = number
	array2 := make([]int, number, number)
	for i := 0; i < number; i++ {
		array2[i] = array1[i]

	}
	// fmt.Println(array2)
	array1 = exchange(array1, number)
	var index int = 0
	if number < len(array1) {
		array1[index] = array1[number]
		index++
		number++
	}
	// fmt.Println(array1)
	index1 := len(array1)
	index2 := len(array2)
	for number2 > 0 {
		index1--
		index2--
		array1[index1] = array2[index2]
		number2--
	}
	fmt.Println("hello", array1)
}
