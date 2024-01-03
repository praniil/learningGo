package dsa

import (
	"fmt"
)

type StacK struct {
	items    []int
	size     int
	capacity int
}

func (s *StacK) NewStack() {
	s.items = make([]int, 0, 5)
	s.size = 0
	s.capacity = 5
}

func (s *StacK) addItems(item int) {
	if s.size == s.capacity {
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
		poppedItem := s.items[s.size-1]
		s.items = s.items[:s.size-1]
		s.size--
		return poppedItem
	}
}

func (s *StacK) DisplayAllItems() {
	for i := s.size - 1; i >= 0; i-- {
		fmt.Println(s.items[i])
	}
}

func TOHoperation(n int, source, aux, dest *StacK) {
	if n > 0 {
		TOHoperation(n-1, source, dest, aux)
		dest.addItems(source.PopItem())
		TOHoperation(n-1, aux, source, dest)
	}
}

func TowerOfHanoi() {
	var stackA, stackB, stackC StacK
	var number int
	fmt.Println("Enter the number of disks for Tower of Hanoi: ")
	stackA.NewStack()
	stackB.NewStack()
	stackC.NewStack()
	fmt.Scanf("%d", &number)
	for i := number; i > 0; i-- {
		stackA.addItems(i)
	}

	TOHoperation(number, &stackA, &stackB, &stackC)

	fmt.Println("Contents of stackC (Destination):")
	stackC.DisplayAllItems()
}
