package main

import "fmt"

func main() {

	// var arrayName [size]elementType

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	fmt.Println("Fruits array:", fruits)

	fmt.Println("Third element:", fruits[2])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100

	fmt.Println("Original array:", originalArray)
	fmt.Println("Copied array:", copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println("Element at index,", i, ":", numbers[i])
	}

	for i, v := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

	// to not use index: _
	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

}
