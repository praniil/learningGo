package arrays

import (
	"fmt"
)

func Palindrome() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 1

	isPalindrome := true
	length := len(a)

	for i := 0; i < length/2; i++ {
		if a[i] != a[length-1-i] {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Println("Palindrome array")
	} else {
		fmt.Println("Not a palindrome array")
	}
}
