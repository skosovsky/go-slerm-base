package main

import "fmt"

func main() {
	fmt.Println(factorialRecursive(10))
	fmt.Println(factorialFor(10))

	limitRecursive()
}

func factorialFor(num int) int {
	factorial := 1

	for i := 1; i <= num; i++ {
		factorial *= i
	}

	return factorial
}

func factorialRecursive(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorialRecursive(num-1)
}

func limitRecursive() {
	for {
		limitRecursive() //nolint: staticcheck
	}
}
