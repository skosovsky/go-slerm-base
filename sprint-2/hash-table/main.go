package main

import "fmt"

const arrSize = 10

type Node struct {
	key   int
	value int
	next  *Node
}

type HashTable struct {
	arr [arrSize]*Node
}

func hash(key int) int {
	return key % arrSize
}

func (h *HashTable) insert(key int, value int) {
	index := hash(key)
	newNode := &Node{
		key:   key,
		value: value,
		next:  nil,
	}

	if h.arr[index] == nil {
		h.arr[index] = newNode
		return
	}

	current := h.arr[index]
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (h *HashTable) get(key int) (int, bool) {
	index := hash(key)

	if h.arr[index] == nil {
		return 0, false
	}

	current := h.arr[index]
	for current != nil {
		if current.key == key {
			return current.value, true
		}
		current = current.next
	}
	return 0, false
}

func main() {
	hashTable := HashTable{}

	hashTable.insert(10, 1)
	hashTable.insert(20, 2)
	hashTable.insert(30, 3)
	hashTable.insert(11, 4)

	val, ok := hashTable.get(10)
	if !ok {
		fmt.Println("Value not found")
	}
	fmt.Println("Value for key", val)

}
