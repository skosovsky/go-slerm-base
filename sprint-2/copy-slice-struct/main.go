package main

import (
	"log"
)

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

	log.Println(houses)

	copyHouses := copySlc(houses)
	log.Println(copyHouses)

	copyHouses[0].name = "My house 3"
	log.Println(copyHouses)
	log.Println(houses)
}

func copySlc(slc []House) []House {
	newSlc := make([]House, 0, len(slc))
	for _, v := range slc {
		newPtr := *v.ptr
		newHouse := House{v.name, &newPtr}
		newSlc = append(newSlc, newHouse)
	}

	return newSlc
}
