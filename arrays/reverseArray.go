package arrays

import(
	"fmt"

)

func ReverseArray(){

	fmt.Println("reversing an array");
	array := [5]int {1, 2, 3, 4, 5};
	for i := 5 ; i > 0 ; i-- {
		fmt.Print(array[i]);
	}

}