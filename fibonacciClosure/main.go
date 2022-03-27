package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prevVal := 1
	prevPrevVal := 0
	firstCall := true

	return func() int {
		if firstCall == true {
			firstCall = false
			return prevPrevVal
		} else {
			result := prevPrevVal + prevVal
			prevPrevVal = prevVal
			prevVal = result

			return result
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
