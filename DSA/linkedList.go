package dsa

import (
	"fmt"
)

type Node struct {
	data int
	next *Node	//pointer to the next node
}
type LinkedLis struct {
	head *Node		//pointer to the head of the list
}

func (li *LinkedLis) Insert(data int) {
	newNode := &Node{ data: data, next: nil }
	if li.head == nil {
		li.head = newNode
		return
	}

	current := li.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (li *LinkedLis) Display () {
	current := li.head
	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
	fmt.Println("Nil")
}

func LinkedList () {
	fmt.Println("IN linked list")
	li := LinkedLis{}
	li.Insert(45);
	li.Insert(30);
	li.Display();
}