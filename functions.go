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

	// Passing a function as an argument
	result1 := applyOperation(1, 10, add)
	fmt.Println(result1)

	// Returning and using a function
	result2 := createMultiplier(2)
	fmt.Println("6*2 =", result2(6))
}

func add(a, b int) int {
	return a + b
}

func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

func createMultiplier(factor int) func (int) int {
	return func (x int) int {
		return x * factor
	}
}
