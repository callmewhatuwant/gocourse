package main

import "fmt"

func main() {

	// func functionName(param1 ... type1) returnType {
	// function body
	// }

	// fmt.Println("Sum of 1,2,3:", sum(1, 2, 3))

	// statement, total := sum("The sum of 1,2,3 is:", 1, 2, 3)
	// fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence2, total2 := sum(2, numbers...)
	fmt.Println("Sequence:", sequence2, "Total:", total2)

}

// func sum(returnString string, nums ...int) (string, int) {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return returnString, total
// }

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
