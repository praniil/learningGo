package arrays

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
    for i:=0; i < len(nums) ; i++{
        for j:=i+1; j< len(nums); j++{
            if nums[i] + nums[j] == 9{
                return []int{i, j}
            }
        }
    }
    return nil
}

func Pointer (){
    nums := []int {2, 11, 7, 15}
    target := 9
    fmt.Println(twoSum(nums, target))
}