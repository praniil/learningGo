package arrays

import (
	"fmt"
)

// interface generic
type Stack struct {
	Size     int
	Capacity int
	Items    []interface{}
}

func NewStack(capacity int) *Stack {
	return &Stack{
		Size:     0,
		Capacity: capacity,
		Items:    make([]interface{}, 0, capacity),
	}
}

func (s *Stack) push(item interface{}) {
	if s.Size < s.Capacity {
		s.Items = append(s.Items, item)
		s.Size++
	} else {
		fmt.Println("the stack is full")
	}
}

func (s *Stack) pop() interface{} {
	if s.Size != 0 {
		// popped := s.Items[s.Size-1]
		s.Items = s.Items[: s.Size-1]
		s.Size--
		// return popped
	}
	fmt.Println("Stack uderflow!")
	return nil
}

func (s *Stack) clear() {
	for  s.Size > 0 {
		s.Items = s.Items[:s.Size-1];			//slice from 0 to size-1
		s.Size--;
	}
}

func (s *Stack) print() {
	fmt.Printf("items available: %d", s.Items...);
}

func StackProperties() {
	myStack := NewStack(5)
	myStack.push(25)
	myStack.push(20)
	myStack.clear();
	myStack.push(20)
	myStack.print();

}
