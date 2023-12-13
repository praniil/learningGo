package arrays

import (
	"fmt"
)

func ArraysProps() {
	fmt.Println("arrays props")
	arrayNumbers := make([]int, 5)
	for i := 0; i < len(arrayNumbers); i++ {
		arrayNumbers[i] = i * 2
	}
	for _, value1 := range arrayNumbers {
		value1 *= 2
		fmt.Println(value1)
	}
	fmt.Println(arrayNumbers)
	for index, value := range arrayNumbers {
		fmt.Println(index, value)
	}
	for _, valueArray := range arrayNumbers {
		fmt.Println(valueArray)
	}
	for i, v := range arrayNumbers {
		if i == 2 {
			v = v * 3
		}
		fmt.Println(v)
	}
}
