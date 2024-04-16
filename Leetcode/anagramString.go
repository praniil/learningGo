package leetcode

import (
	"fmt"
)

// func swap(array []rune){
// for i := 0; i < len(array)-1; i++ {
// for j := i + 1; j < len(array); j++ {
// if array[i] > array[j] {
// temp := array[i]
// array[i] = array[j]
// array[j] = temp
// }
// }
// }
// }
//
// func isAnagram(s string, t string) bool {
// rune1 := []rune(strings.ToLower(s))
// rune2 := []rune(strings.ToLower(t))
// if (len(rune1) != len(rune2)) {
// return false
// }
// swap(rune1)
// swap(rune2)
// for i := 0; i < len(rune1); i++ {
// if rune1[i] != rune2[i] {
// return false
// }
// }
// return true
// }

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}
	for _, char := range t {
		count[char]--
		if count[char] < 0 {
			return false
		}
	}
	return true
}

func AnagramString() {
	string1 := "pranil"
	string2 := "ranpil"
	result := isAnagram(string1, string2)
	fmt.Println(result)
}
