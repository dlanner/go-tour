package main

import "golang.org/x/tour/pic"

type fn func(int, int) uint8

func PicClosure(f fn) func(int, int) [][]uint8 {
	return func(dx, dy int) [][]uint8 {
		var matrix [][]uint8
		for x := 0; x < dx; x++ {
			var ys []uint8
			for y := 0; y < dy; y++ {
				newY := f(x, y)
				ys = append(ys, uint8(newY))
			}
			matrix = append(matrix, ys)
		}
		return matrix
	}
}

func f1(x, y int) uint8 {
	return uint8((x + y) / 2)
}

func f2(x, y int) uint8 {
	return uint8(x * y)
}

func f3(x, y int) uint8 {
	return uint8(x ^ y)
}

func main() {
	pic.Show(PicClosure(f1))
	// pic.Show(PicClosure(f2))
	// pic.Show(PicClosure(f3))
}
