package leetcode

import(
	"fmt"
)

func twoSum(nums []int, target int) []int {
    arr := make([]int, 2)
	for i:= 0; i < len(nums) - 1; i++{
		for j := 0; j < len(nums); j++ {
			if nums[i] + nums[j] == target{
				arr[0] = i
				arr[1] = j
			}
		}
	}
	return []int{arr[0], arr[1]}
}
func TwoSum() {
	fmt.Println("in two sum")
	arr := []int {2, 7, 11, 15}
	result := make([]int, 2)
	var target int = 9
	result[0] = twoSum(arr, target)[0]
	result[1] = twoSum(arr, target)[1]
	fmt.Println(result)
}