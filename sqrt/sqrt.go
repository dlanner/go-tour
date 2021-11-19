package main

import (
	"fmt"
)

func contains(floats []float64, fval float64) bool {
	for _, v := range floats {
		if v == fval {
			return true
		}
	}

	return false
}

func Sqrt(x float64) float64 {
	var seenValues []float64
	z := 1.0
	seenValues = append(seenValues, z)

	//for !contains(seenValues, z) {
	for i := 0; i < 100; i++ {
		z -= (z*z - x) / (2 * z)
		if contains(seenValues, z) {
			fmt.Printf("Sqrt(%v): %v (iterations: %v)\n", x, z, len(seenValues))
			return z
		}
		seenValues = append(seenValues, z)
	}
	return z
}

func main() {
	for i := 2.0; i < 100.0; i++ {
		Sqrt(i)
	}
}
