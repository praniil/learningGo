package arrays

import (
	"fmt"
)

func Array() {
	var a [3]int
	var a1 [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	for i := 0; i < len(a); i++ {
		a1[i] = a[i] * a[i]
		fmt.Println(a1[i])
	}
}
