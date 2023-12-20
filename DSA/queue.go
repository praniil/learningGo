package dsa

import (
	"fmt"
)

type QueStr struct {
	items    []int
	size     int
	capacity int
	rear     int
	front    int
}

func (qs *QueStr) AddItems(item int) {
	if qs.rear == qs.capacity-1 {
		fmt.Println("queue is full")
	} else {
		qs.rear = qs.rear + 1
		qs.items = append(qs.items, item)
		fmt.Println(qs.items)
		qs.size++
	}
}

func (qs *QueStr) WithdrawItems() {
	if qs.front > qs.rear {
		fmt.Println("queue is empty")
	} else {
		fmt.Printf("%d is dequeued", qs.items[qs.front])
		qs.front++
	}
}

func (qs *QueStr) DisplayAllElements() {
	for i := qs.front; i <= qs.rear; i++ {
		fmt.Println(qs.items[i])
	}
}

func Queue() {
	queue := QueStr{items: make([]int, 0, 5), size: 0, capacity: 5, rear: -1, front: 0}
	queue.AddItems(20)
	queue.AddItems(24)
	queue.AddItems(25)
	queue.AddItems(23)
	queue.AddItems(22)
	queue.WithdrawItems()
	queue.DisplayAllElements()
}
