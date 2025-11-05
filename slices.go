package main

import "fmt"

func main() {

	// var sliceName [] ElementType

	// var numbers []int
	// var numbers1 = []int{1,2,3}

	// numbers2 = []int{9,8,7}

	a := [5]int{1, 2, 3, 4, 5}
	slice1 := a[1:4]

	fmt.Println(slice1)

	slice1 = append(slice1, 6, 7)
	fmt.Println(slice1)

	sliceCopy := make([]int, len(slice1))
	copy(sliceCopy, slice1)
	fmt.Println("Copy:", sliceCopy)

	// var nilSlice []int

	for i, v := range slice1 {
		fmt.Println(i, v)
	}
	fmt.Println("Element at index 3:", slice1[3])

	slice1[3] = 50
	fmt.Println("Element at index 3:", slice1[3])

}
