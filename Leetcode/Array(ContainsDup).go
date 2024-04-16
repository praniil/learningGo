package leetcode

import (
	"fmt"
)

func containsDuplicate(nums []int) bool{
	seen := make(map[int]bool)
	for _, num:= range nums{
		if seen[num] {		//if true
			return true
		}
		seen[num] = true
	}
	return false
}

func ContDup() {
	fmt.Println("in contDUp")
	// var array = [...]int {1, 1, 1, 3 ,3 ,4 ,3 ,2 ,4 ,2}	
	var array = [...]int {1, 2, 3, 4, 4}	
	fmt.Println(array)
	x:= containsDuplicate(array[:])
	fmt.Println(x)
	fmt.Println(array)
}
