package arrays

import(
	"fmt"
)

func Factorial (n int) int {
	if n == 0{
		return 1
	}
	fact := n * Factorial(n - 1)
	return fact
}


func Recursion (){
	number := 5
	factorial_of_number := Factorial(number);
	fmt.Println(factorial_of_number)
}

