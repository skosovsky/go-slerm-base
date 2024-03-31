package main

import (
	"fmt"
	"sync"
)

// sumChannels
// in - слайс входных каналов, в которые приходят числа
// Признак окончания данных в канале - канал закрыт
func sumChannels(inputs []chan int64) int64 {
	var wg sync.WaitGroup

	out := make(chan int64, len(inputs))

	// Read from inputs ([]chan)
	for _, input := range inputs {
		var partSum int64
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range input {
				partSum += num
			}
			out <- partSum
		}()
	}

	wg.Wait()
	close(out)

	// Read from out
	var result int64
	for num := range out {
		result += num
	}
	return result
}

func main() {
	countChan := 1000
	chs := make([]chan int64, 0, countChan)

	for range 100 {
		var ch = make(chan int64, countChan)
		for num := range 1000 {
			ch <- int64(num)
		}
		close(ch)
		chs = append(chs, ch)
	}

	result := sumChannels(chs)
	fmt.Println(result) //nolint:forbidigo // it's learning code

	// check fix input in 1.22
	inputs := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}, {5, 5, 5}}
	var sumAll int
	var wg sync.WaitGroup
	for _, input := range inputs {
		var sum int
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, num := range input { // it's fix in 1.22
				sum += num
			}
			sumAll += sum
		}()
	}
	wg.Wait()
	fmt.Println(sumAll) //nolint:forbidigo // it's learning code
}
