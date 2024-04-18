package main

import (
	"log"
	"strings"
	"sync"
)

func mapRace() {
	var data = make(map[int]int)

	var wg sync.WaitGroup
	wg.Add(2) //nolint:gomnd // it's learning code
	go func() {
		defer wg.Done()
		for i := range 10 {
			data[i] = i
		}
	}()
	go func() {
		defer wg.Done()
		for i := range 10 {
			data[i+100] = i + 100 //nolint:gomnd // it's learning code
		}
	}()
	wg.Wait()
	log.Println(len(data))
}

func countWords(line string) {
	var data = make(map[string]int)
	var wg sync.WaitGroup

	words := strings.Split(line, " ")
	for _, word := range words {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data[word]++
		}()
	}
	defer wg.Wait()

	log.Println(len(data))
}

func countWordsMutex(line string) {
	var data = make(map[string]int)
	var wg sync.WaitGroup
	var m sync.Mutex

	words := strings.Split(line, " ")
	for _, word := range words {
		wg.Add(1)
		m.Lock()
		go func() {
			defer wg.Done()
			defer m.Unlock()
			data[word]++
		}()
	}
	wg.Wait()

	log.Println(data["да"])
	log.Println(len(data))
}

func countWordsSyncMap(line string) {
	var data sync.Map
	var wg sync.WaitGroup

	words := strings.Split(line, " ")
	for _, word := range words {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if iNum, isLoad := data.Load(word); isLoad {
				num, ok := iNum.(int)
				if ok {
					num++
					data.Store(word, num)
					return
				}
				return
			}
			data.Store(word, 1)
		}()
	}
	wg.Wait()

	log.Println(data.Load("да"))
}

func race() {
	inputs := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}, {5, 5, 5}}
	var sumAll int
	var wg sync.WaitGroup

	for _, input := range inputs {
		var sum int
		wg.Add(1)
		go func(input []int) {
			defer wg.Done()
			for _, num := range input {
				sum += num
			}
		}(input)

		wg.Wait()
		sumAll += sum
	}
	log.Println(sumAll)
}

func raceTwo() {
	inputs := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}, {4, 4, 4}, {5, 5, 5}}
	var sumAll int
	var wg sync.WaitGroup
	for _, input := range inputs {
		var sum int
		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, num := range input {
				sum += num
			}
			sumAll += sum
		}()
	}
	wg.Wait()
	log.Println(sumAll)
}

func main() {
	mapRace()                                                                                 // fatal error: concurrent map writes
	countWords("ах если бы да ка бы да во рту выросли грабы да был бы не рот а целый огород") // fatal error: concurrent map writes
	countWordsMutex("ах если бы да ка бы да во рту выросли грабы да был бы не рот а целый огород")
	countWordsSyncMap("ах если бы да ка бы да во рту выросли грабы да был бы не рот а целый огород")
	race()
	raceTwo()
}
