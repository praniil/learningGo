// package dsa

// import (
// 	"fmt"
// )

// type QueStr struct {
// 	items    []int
// 	size     int
// 	capacity int
// 	rear     int
// 	front    int
// }

// func (qs *QueStr) AddItems(item int) {
// 	if qs.rear == qs.capacity-1 {
// 		fmt.Println("queue is full")
// 	} else {
// 		qs.rear = qs.rear + 1
// 		qs.items = append(qs.items, item)
// 		fmt.Println(qs.items)
// 		qs.size++
// 	}
// }

// func (qs *QueStr) WithdrawItems() {
// 	if qs.front > qs.rear {
// 		fmt.Println("queue is empty")
// 	} else {
// 		fmt.Printf("%d is dequeued", qs.items[qs.front])
// 		qs.front++
// 	}
// }

// func (qs *QueStr) DisplayAllElements() {
// 	for i := qs.front; i <= qs.rear; i++ {
// 		fmt.Println(qs.items[i])
// 	}
// }

// func Queue() {
// 	queue := QueStr{items: make([]int, 0, 5), size: 0, capacity: 5, rear: -1, front: 0}
// 	queue.AddItems(20)
// 	queue.AddItems(24)
// 	queue.AddItems(25)
// 	queue.AddItems(23)
// 	queue.AddItems(22)
// 	queue.WithdrawItems()
// 	queue.WithdrawItems()
// 	queue.WithdrawItems()
// 	queue.WithdrawItems()
// 	queue.WithdrawItems()
// 	queue.WithdrawItems()
// 	queue.DisplayAllElements()
// }

package dsa

import (
	"fmt"
)

var array [20]int
var front int = 0
var rear int = -1

func Enqueue() {
	if front < rear {
		fmt.Println("queue is full")
	} else {
		var item int
		fmt.Println("enter the value of item: ")
		fmt.Scanf("%d", item)
		rear++
		array[rear] = item
	}
}

func Dequeue() {
	if front > rear {
		fmt.Println("queue is empty")
	} else {
		fmt.Println("the item dequeued is:", array[front])
		front++
	}
}
func DisplayFront() {
	if front > rear {
		fmt.Println("the queue is empty")
	} else {
		fmt.Println(array[front])
	}
}
func DisplayAll() {
	if front > rear {
		fmt.Println("the queue is empty")
	} else {
		for i := front; i <= rear; i++ {
			fmt.Printf("%d", array[i])
		}
	}
}

func Queue() {
	var number int = 1
	for number == 1 {
		fmt.Println("1.Enqueue 2.Dequeue 3.Display the first Element 4.Display all 5.Exit")
		var choice int
		fmt.Println("Enter the choice")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			Enqueue()
			break
		case 2:
			Dequeue()
			break
		case 3:
			DisplayFront()
		case 4:
			DisplayAll()
		case 5:
			number = 0
		default:
			fmt.Println("wrong choice")

		}
	}
}
