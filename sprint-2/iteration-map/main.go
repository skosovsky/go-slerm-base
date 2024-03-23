package main

import "fmt"

func main() {
	data := map[int]int{1: 1, 2: 2, 3: 3, 4: 4}
	iterationByMap(data)
	fmt.Println(data)
}

func iterationByMap(data map[int]int) {
	for key, value := range data {
		fmt.Printf("%d:%d ", key, value)
	}
	fmt.Println()
}
