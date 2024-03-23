package main

import "fmt"

type House struct {
	name string
	ptr  *string
}

func main() {
	house1 := House{
		name: "My House",
		ptr:  new(string),
	}
	house2 := House{
		name: "My House 2",
		ptr:  new(string),
	}

	word := "word"
	house1.ptr = &word

	var houses []House
	houses = append(houses, house1, house2)

	fmt.Println(houses)

	copyHouses := copySlc(houses)
	fmt.Println(copyHouses)

	copyHouses[0].name = "My house 3"
	fmt.Println(copyHouses)
	fmt.Println(houses)
}

func copySlc(slc []House) []House {
	// TODO переделать функцию, чтобы пересоздавала вручную все сущности
	newSlc := make([]House, len(slc))
	copy(newSlc, slc)

	return newSlc
}
