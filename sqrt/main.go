package main

import (
	"fmt"
	"math"
)

const (
	MAX_ITRS       = 10
	STARTING_VALUE = 1.0
	DELTA_ITR      = 0.0000000001
)

func Sqrt(x float64) float64 {
	z := STARTING_VALUE
	prevZ := 0.0

	for i := 0; ; i++ {
		prevZ = z
		z -= (z*z - x) / (2 * z)
		fmt.Println(i, z)

		if math.Abs(z-prevZ) < DELTA_ITR {
			break
		}
	}

	return z
}

func main() {
	value := 102314.0
	fmt.Printf("Sqrt of %v is %v\n", value, Sqrt(value))

	fmt.Printf("Sqrt of %v is %v\n", value, math.Sqrt(value))
}
