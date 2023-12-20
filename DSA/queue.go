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
	if qs.rear >= qs.capacity {
		fmt.Println("queue is full")
	} else {
		qs.rear++
		qs.items[qs.rear] = item
		qs.size++
	}
}

func (qs *QueStr) WithdrawItems() {
	
}

func Queue() {
	queue := QueStr{items: make([]int, 0, 5), size: 0, capacity: 5, rear: -1, front: 0}
	queue.AddItems(20)
	queue.AddItems(24)
	queue.AddItems(25)
	queue.AddItems(23)
	queue.AddItems(22)
}
