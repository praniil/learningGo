//multiple return value

package arrays

import (
	"fmt"
)

func sumProduct() (int, int) {
	a := 6
	b := 5
	var sum int = a + b
	var product int = a * b
	return sum, product
}

func Functions() {
	
	fmt.Println(sumProduct())
}
