package leetcode

import (
	"fmt"
)

func operation(arr [4]int, target int) []int {
	var sum []int
	for i := 0; i < len(arr); i++ {
		sum = append(sum, i)
		for j := i + 1; j < len(arr); j++ {
			if arr[j] == target-arr[i] {
				sum = append(sum, j)
				return sum
			} else if arr[j]+arr[j+1] == target-arr[i] {
				sum = append(sum, j, j+1)
				return sum
			} else {
				sum = append(sum, j, j+1, j+2)
				return sum
			}
		}
	}
	return nil
}

func TwoSum() {
	arr := [4]int{2, 3, 4, 15}
	target := 9
	arrResult := operation(arr, target)
	fmt.Println(arrResult)

}
