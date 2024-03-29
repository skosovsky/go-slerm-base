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
	var chs []chan int64

	for range 100 {
		var ch = make(chan int64, 1000)
		for num := range 1000 {
			ch <- int64(num)
		}
		close(ch)
		chs = append(chs, ch)
	}

	result := sumChannels(chs)
	fmt.Println(result)
}
