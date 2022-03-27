package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello " + name)
}

func main() {
	sayHello("Ciprian")
	i := 0

	for i < 10 {
		fmt.Println(i)
		i++
	}
}
