package arrays

import (
	"fmt"
)

type Base struct{
	number int
}

func (b *Base) describe() string {
	return fmt.Sprintf("base with number = %v", b.number)
}

type Container struct {
	Base
	str string
}
//while making a method if we use pointer and later on copy the object of that struct to another object of the same struct, same address ma data members haru basxa


func StructEmbedding () {
	base:= Base{
		number: 45,

	}

	base1 := &base
	fmt.Println("base1.number: ", &base1.number)
	fmt.Println("base.number: ", &base.number)

	co := Container{
		Base: Base{
			number :1,
		},
		str : "Pranil",
	}

	fmt.Printf("co= {number: %v, string: %v} \n", co.number, co.str)
	fmt.Println("alos num: ", co.Base.number)
	fmt.Println(co.describe())
	
	type Desriber interface {
		describe() string
	}

	var d Desriber = &co
	fmt.Println("describer: ", d.describe())
	

}