package dsa

import (
	"fmt"
)

type Stack struct {
	items    []string
	size     int
	capacity int
}

func (s *Stack) Push(item string) {
	if s.size < s.capacity {
		s.items = append(s.items, item)
	} else {
		fmt.Println("overflow")
	}
}

func (s *Stack) Pop(item string) string {
	if s.size > 0 {
		s.items = s.items[:s.size-1]
		return s.items[s.size]
	}
	return ""
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
	operatorStack := Stack{items: make([]string, 0), size: 0, capacity: 25}
	fmt.Println(operatorStack)
	postfixStack := Stack{items: make([]string, 0), size: 0, capacity: 25}
	fmt.Println(postfixStack)

}
