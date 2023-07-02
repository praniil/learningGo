package arrays

import(
	"fmt"
)

func incrementer() func() int{
	count := 0

	increment := func() int {
		count = count + 1
		return count
	}

	return increment
}

func Closures(){
	counter := incrementer()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}