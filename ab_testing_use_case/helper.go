package viant

import (
	"math/rand"
	"math"
)

func randInt(min, max int) int {
	return min + int( math.Round(float64(max - min) * float64(rand.Intn(10000)) / 10000.0))
}

func randFloat(min, max float64) float64 {
	return min + (max - min) * float64(rand.Intn(10000)) / 10000.0
}
