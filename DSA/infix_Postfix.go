package dsa

import (
	"fmt"
)

type Stack struct {
	items    []rune
	size     int
	capacity int
}

func (s *Stack) Push(item rune) {
	if s.size < s.capacity {
		s.items = append(s.items, item)
		s.size++
	} else {
		fmt.Println("overflow")
	}
}

func (s *Stack) Pop() rune {
	if s.size > 0 {
		poppedItem := s.items[s.size-1]
		s.items = s.items[:s.size-1]
		s.size--
		return poppedItem
	}
	return 0
}

func (s *Stack) Peek() rune {
	if s.size > 0 {
		return rune(s.items[s.size-1])
	}
	return 0
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func Precedence(operator rune) int {
	switch operator {
	case '^', '$':
		return 3
	case '/', '*':
		return 2
	case '+', '-':
		return 1
	}
	return 0
}

func InfixPostfix() {
	operatorStack := Stack{items: make([]rune, 0), size: 0, capacity: 25}
	postfixStack := Stack{items: make([]rune, 0), size: 0, capacity: 25}
	expression := "A+(B*C-(D/E$F)*G)*H"
	for _, char := range expression {
		if char >= 'A' && char <= 'Z' {
			postfixStack.Push(char)
		} else if char == '(' {
			operatorStack.Push(char)
		} else if char == ')' {
			for operatorStack.Peek() != '(' && !operatorStack.isEmpty() {
				opPop := operatorStack.Pop()
				postfixStack.Push(opPop)
			}
			if operatorStack.Peek() == '(' {
				operatorStack.Pop()
			}
		} else {
			for !operatorStack.isEmpty() && Precedence(char) <= Precedence(operatorStack.Peek()) {
				postfixStack.Push(operatorStack.Pop())
			}
			operatorStack.Push(char)
		}
	}
	for !operatorStack.isEmpty() {
		postfixStack.Push(operatorStack.Pop())
	}
	fmt.Println(string(postfixStack.items))
}
