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
	//this is a map named count which has rune as a key and int as value of that key, default value of key is 0 for int
	count := make(map[rune]int)
	//initially for each key(here it is rune) by default the value is 0
	//for each encounter of a char, that (key= char), value= value+1 char's value increases by 1
	for _, char := range s {
		count[char]++
	}
	//if the char in t is present in map, it obviously has the value in the map greater than zero,
	//for each encounter of the char in t the count map's char's key value is decrease by 1
	//if the value goes negative the char in not present in the initial string because be didnt find the balance while decreaseing the value
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
