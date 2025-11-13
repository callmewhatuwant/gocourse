package main

import "fmt"

func main() {

	sequence := adder()
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	susbtractor := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(susbtractor(1))
	fmt.Println(susbtractor(2))
	fmt.Println(susbtractor(3))
	fmt.Println(susbtractor(4))
	fmt.Println(susbtractor(5))

}

func adder() func() int {
	i := 0
	fmt.Println("previous value of i:", i)
	return func() int {
		i++
		fmt.Println("added 1 to i")
		return i
	}
}
