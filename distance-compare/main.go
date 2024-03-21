package main

import (
	"fmt"
	"math"
)

func main() {
	var lat1, lon1 float32 = 55.873733, 37.401319
	var lat2, lon2 float32 = 55.873733, 37.401319
	fmt.Println(distanceCompare(lat1, lon1, lat2, lon2))
}

func distanceCompare(lat1, lon1 float32, lat2, lon2 float32) bool {
	return math.Abs(float64(lat1)) == math.Abs(float64(lat2)) && math.Abs(float64(lon1)) == math.Abs(float64(lon2))
}
