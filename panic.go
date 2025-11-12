package main

import "fmt"

func main() {

	// panic(interface{})

	// Example of valid input
	process(10)

	// Example of invalid input
	process(-3)

}

func process(input int) {
	if input < 0 {
		panic("input must be non-negative number")
	}
	fmt.Println("Processing input:", input)
}
