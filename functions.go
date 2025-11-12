package main

import "fmt"

func main() {

	// funx <name> (parameters list) returnType {
	// code block
	// return value
	// }

	//sum := add(1, 2)
	fmt.Println(add(2, 3))

	// func ()  {
	// 	fmt.Println("Anonymous Function")
	// }()

	greet := func() {
		fmt.Println("Anonymous Function")
	}

	greet()

	operation := add

	result := operation(3, 5)

	fmt.Println(result)

}

func add(a, b int) int {
	return a + b
}
