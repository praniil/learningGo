package arrays

import (
	"fmt"
)

func sum(nums ...int) int { // variable number of arguments
	sum := 0
	for _, num := range nums{
		sum = sum + num
	}
	return sum
}

func Variadic_function (){
	total_sum := sum(1, 2, 3, 4, 5)
	fmt.Println(total_sum)

	nums := []int {1, 2, 3, 40}
	total_sum_array := sum(nums...)
	fmt.Println(total_sum_array)
}
