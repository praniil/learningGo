package dsa

import (
	"fmt"
)

type StacK struct {
	items    []int
	size     int
	capacity int
}

func (s *StacK) NewStack() *StacK {
	stack := &StacK{items: make([]int, 0, 5), size: 0, capacity: 5}
	return stack
}

func (s *StacK) addItems(item int) {
	if s.size == s.capacity-1 {
		fmt.Println("stack is full")
	} else {
		s.items = append(s.items, item)
		s.size++
	}
}

func (s *StacK) PopItem() int {
	if s.size == 0 {
		fmt.Println("the stack is empty")
		return 0
	} else {
		items := s.items[:s.size-1]
		poppedItem := s.items[s.size-1]
		s.items = items
		s.size--
		return poppedItem
	}
}

func (s *StacK) DisplayAllItems() {
	for i := 0; i < s.size; i++ {
		fmt.Println(s.items[i])
	}
}

func TOHoperation(number int, sA *StacK, sB *StacK, sC *StacK) {
	if number > 0 {

		TOHoperation(number-1, sA, sC, sB)
		sC.addItems(sA.PopItem())
		TOHoperation(number-1, sC, sA, sB)
	}
}

func TowerOfHanoi() {
	var stackA, stackB, stackC StacK
	var number int
	fmt.Println("enter the number of hanoi you want: ")
	stackA.NewStack()
	stackB.NewStack()
	stackC.NewStack()
	fmt.Scanf("%d", &number)
	for i := 1; i <= number; i++ {
		stackA.addItems(i)
	}
	TOHoperation(number, &stackA, &stackB, &stackC)
	stackA.DisplayAllItems()
	fmt.Println("-------")
	stackB.DisplayAllItems()
	fmt.Println("-------")
	stackC.DisplayAllItems()
}
