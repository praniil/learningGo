//atruct method

package arrays

import "fmt"

type Employee struct{
	name string
	phoneNumber int 
}

func (emp1 *Employee) set_Name(name string) {
	emp1.name = name
}

func (emp1 Employee) set_phone(phoneNo int){
	emp1.phoneNumber = phoneNo
}

func (emp1 *Employee) get_Name() string {
	return emp1.name;
}

func (emp1 Employee) get_Phone() int{
	return emp1.phoneNumber
}

func Struct(){
	emp1 := Employee{name: "", phoneNumber: 0}
	emp1.set_Name("Pranil")
	fmt.Println(emp1.get_Name())
}

