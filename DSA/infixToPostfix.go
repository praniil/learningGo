package dsa

import (
	"fmt"
)

type Stack struct {
	Size     int
	Capacity int
	Items    []string
}

// its like a constructor
func NewStack(capacity int) *Stack {
	return &Stack{
		Size:     0,
		Capacity: capacity,
		Items:    make([]string, 0, capacity),
	}
}

func (s *Stack) Push(item string) string {
	if s.Size < s.Capacity {
		pushedItem := s.Items[s.Size+1]
		s.Items = append(s.Items, item)
		s.Size++
		return pushedItem
	} else {
		fmt.Println("the stack is kinda full")
		return ""
	}
}

func (s *Stack) Pop(item string) string {
	if s.Size > 0 {
		poppedItem := s.Items[s.Size-1]
		s.Items = s.Items[:s.Size-1]
		s.Size--
		return poppedItem
	} else {
		fmt.Println("the stack is kinda empty")
		return ""
	}
}

func (s *Stack) DisplayStack() {
	fmt.Println(s.Items)
}

func (s *Stack) LastItem() string {
	fmt.Println(s.Items[s.Size])
	return s.Items[s.Size]
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '/', '*':
		return 2
	case '^', '$':
		return 3
	}
	return -1
}

func InfinixToPostfix() {
	var expression string = "A+(B*C-(D/E$F)*G)*H"

	stackPostFixString := NewStack(20)
	stack := NewStack(20)
	for i := 0; i < len(expression); i++ {
		char := rune(expression[i])
		if ( char >= 'A' && char <='Z') || ( char >= 'a' && char <= 'z') {
			stackPostFixString.Push(string(char))
		} else if char == '(' {
			stack.Push(string(char))
		} else if char == ')' {
			for stack.LastItem()[0] != '(' {
				stack.Push(string(char))
			}
			stack.Pop("")
		}

	}
	stackPostFixString.DisplayStack()
	stack.DisplayStack()
}
