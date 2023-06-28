package arrays

import "fmt"

func Using_arrays(){
	// var number[5] int;
	// fmt.Print("emp:", number);
	number := [5] int{1, 2, 3, 4, 5}
	number[4] = 100;
	// fmt.Println("set: ", number);
	fmt.Print(number);
	
}