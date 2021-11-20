package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func contains(floats []float64, fval float64) bool {
	for _, v := range floats {
		if v == fval {
			return true
		}
	}

	return false
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		var seenValues []float64
		z := 1.0
		seenValues = append(seenValues, z)

		//for !contains(seenValues, z) {
		for i := 0; i < 100; i++ {
			z -= (z*z - x) / (2 * z)
			if contains(seenValues, z) {
				break
			}
			seenValues = append(seenValues, z)
		}
		return z, nil
	}
}

func main() {
	for i := -5.0; i < 5.0; i++ {
		sqrt, err := Sqrt(i)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Sqrt(%v): %v\n", i, sqrt)
		}
	}
}
