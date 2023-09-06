package arrays

import "fmt"

func BinarySearch (){
	array := []int {1, 2, 3, 4, 5}
	low := 0;
	high := len(array) - 1
	search := 5
	for low <= high {
		middle := (low + high) / 2
		guess := middle;
		if (search == array[guess]){
			fmt.Println("found in ", guess, "index")
			break
		}else if (search <= guess){
			high = guess - 1 
		}else{
			low = guess + 1
		}
	}
}