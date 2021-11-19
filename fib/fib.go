package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prevVal := 0
	prevPrevVal := 0
	currentVal := 0
	return func() int {
		if currentVal == 0 {
			prevPrevVal = prevVal
			prevVal = currentVal
			currentVal = 1
			return currentVal
		} else if currentVal == 1 {
			prevPrevVal = prevVal
			prevVal = currentVal
			currentVal = 2
			return currentVal
		} else {
			prevPrevVal = prevVal
			prevVal = currentVal
			currentVal = prevPrevVal + prevVal
			return currentVal
		}
	}
}

func printVals(prevPrevVal, prevVal, currentVal int) {
	fmt.Printf("n-2: %v, n-1: %v, n: %v\n", prevPrevVal, prevVal, currentVal)
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
