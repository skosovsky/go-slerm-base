package main

import (
	"encoding/json"
	"log"
	"os"
)

type SalePoint struct {
	ID      int    `json:"id"`
	IP      string `json:"IP"`
	Comment string `json:"-"`
}

func main() {
	fileConfig, err := os.ReadFile("sale_points.json")
	if err != nil {
		log.Println(err)
		return
	}

	var salePoints []SalePoint
	err = json.Unmarshal(fileConfig, &salePoints)
	if err != nil {
		log.Println(err)
		return
	}

	salePointsIP := make(map[int]string, len(salePoints))
	for _, salePoint := range salePoints {
		salePointsIP[salePoint.ID] = salePoint.IP
	}

	log.Println(salePointsIP)
}
