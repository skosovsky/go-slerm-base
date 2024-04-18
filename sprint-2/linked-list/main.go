package main

import (
	"log"
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
		log.Printf("%d - > ", current.value)
		current = current.next
	}
	log.Println("nil")
}

func reverseLinkedList(list *LinkedList) *LinkedList {
	var reversedHead *Node
	current := list.head

	for current != nil {
		newNode := &Node{
			value: current.value,
			next:  nil,
		}
		newNode.next = reversedHead
		reversedHead = newNode
		current = current.next
	}

	return &LinkedList{head: reversedHead}
}

func main() {
	list := LinkedList{
		head: nil,
	}
	list.append(1)
	list.append(2) //nolint:gomnd // it's learning code
	list.append(3) //nolint:gomnd // it's learning code
	list.append(4) //nolint:gomnd // it's learning code
	log.Println("Original Linked List:")
	list.print()

	reversedList := reverseLinkedList(&list)
	log.Println("Reversed Linked List:")
	reversedList.print()

	log.Println("Original Linked List (unchanged):")
	list.print()
}
