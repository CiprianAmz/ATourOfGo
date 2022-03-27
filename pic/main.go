package main

import "golang.org/x/tour/pic"

func picValGenerator(x, y uint8) uint8 {
	return (x + y) / 2
}

func Pic(dx, dy int) (result [][]uint8) {
	result = make([][]uint8, dy)

	for i := 0; i < len(result); i++ {
		result[i] = make([]uint8, dx)

		for j := 0; j < len(result[i]); j++ {
			result[i][j] = picValGenerator(uint8(i), uint8(j))
		}
	}

	return
}

func main() {
	pic.Show(Pic)
}
