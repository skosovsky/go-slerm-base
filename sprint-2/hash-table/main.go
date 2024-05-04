package main

import (
	"log"
)

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
	hashTable := HashTable{
		arr: [10]*Node{},
	}

	hashTable.insert(10, 1) //nolint:mnd // it's learning code
	hashTable.insert(20, 2) //nolint:mnd // it's learning code
	hashTable.insert(30, 3) //nolint:mnd // it's learning code
	hashTable.insert(11, 4) //nolint:mnd // it's learning code

	val, ok := hashTable.get(10) //nolint:mnd // it's learning code
	if !ok {
		log.Println("Value not found")
	}
	log.Println("Value for key", val)
}
