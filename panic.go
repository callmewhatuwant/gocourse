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

	defer fmt.Println("Defferred 1")
	defer fmt.Println("Defferred 2")

	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be non-negative number")
	}
	fmt.Println("Processing input:", input)
}
