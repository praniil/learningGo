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
		fmt.Print(s.Items...)
	} else {
		fmt.Println("the stack is full")
	}
}

func (s *Stack) pop() interface{} {
	if s.Size != 0 {
		// popped := s.Items[s.Size-1]
		s.Items = s.Items[: s.Size-1]
		s.Size--
		fmt.Print(s.Items...)
		// return popped
	}
	fmt.Println("Stack uderflow!")
	return nil
}

func StackProperties() {
	myStack := NewStack(5)
	myStack.push(25)
	myStack.push(20)
	myStack.pop()
	myStack.pop()
}
