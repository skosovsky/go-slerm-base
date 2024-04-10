package main

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (l *LinkedList) append(value int) {
	newNode := &Node{value: value, next: nil}
	if l.head == nil {
		l.head = newNode
		return
	}

	lastNode := l.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = newNode
}

func (l *LinkedList) print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d - > ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func reverseLinkedList(list *LinkedList) *LinkedList {
	var reversedHead *Node
	current := list.head

	for current != nil {
		newNode := &Node{value: current.value}
		newNode.next = reversedHead
		reversedHead = newNode
		current = current.next
	}

	return &LinkedList{head: reversedHead}
}

func main() {
	list := LinkedList{}
	list.append(1)
	list.append(2)
	list.append(3)
	list.append(4)
	fmt.Println("Original Linked List:")
	list.print()

	reversedList := reverseLinkedList(&list)
	fmt.Println("Reversed Linked List:")
	reversedList.print()

	fmt.Println("Original Linked List (unchanged):")
	list.print()
}
