package main

import (
	"fmt"
	"math"
)

func main() {
	var lat1, lon1 float32 = 55.873733, 37.401319
	var lat2, lon2 float32 = 55.870035, 37.408641
	fmt.Println(distance(lat1, lon1, lat2, lon2))
}

func distance(lat1, lon1 float32, lat2, lon2 float32) float32 {
	const earthRadius = 6371

	lat1Rad := lat1 * (math.Pi / 180)
	lon1Rad := lon1 * (math.Pi / 180)
	lat2Rad := lat2 * (math.Pi / 180)
	lon2Rad := lon2 * (math.Pi / 180)

	deltaLat := lat2Rad - lat1Rad
	deltaLon := lon2Rad - lon1Rad

	a := math.Pow(math.Sin(float64(deltaLat)/2), 2) + math.Cos(float64(lat1Rad))*math.Cos(float64(lat2Rad))*math.Pow(math.Sin(float64(deltaLon)/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return float32(earthRadius * c)
}

// 55.873733, 37.401319
// 55.870035, 37.408641
