package dsa

import (
	"fmt"
)

type Node struct {
	data int
	next *Node //pointer to the next node
}
type LinkedLis struct {
	head *Node //pointer to the head of the list
}

func (li *LinkedLis) append(data int) {
	newNode := &Node{data: data, next: nil}
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

func (li *LinkedLis) insertNext(data int, prevNode *Node) {
	if prevNode == nil {
		fmt.Println("Previous node cannot be nil")
	}
	newNode := &Node{data: data, next: nil}
	newNode.next = prevNode.next
	prevNode.next = newNode
}

func (li *LinkedLis) prepend(data int, nextNode *Node) {
	if nextNode == nil {
		fmt.Println("Next node cannot be nil")
	}
	newNode := &Node{data: data, next: nil}
	newNode.next = li.head
	li.head = newNode

}

func (li *LinkedLis) Display() {
	current := li.head
	for current != nil {
		fmt.Printf("%d ->", current.data)
		current = current.next
	}
	fmt.Println("Nil")
}

func LinkedList() {
	fmt.Println("IN linked list")
	li := LinkedLis{}
	li.append(45)
	li.append(30)
	li.append(90)

	currentNode := li.head
	for currentNode != nil && currentNode.data != 90 {
		currentNode = currentNode.next
	}
	if currentNode != nil {
		li.prepend(200, currentNode)
		li.insertNext(2, currentNode)
	}
	li.Display()
}
