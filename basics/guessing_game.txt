package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to guessing game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess the number")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		// %  to store at address of var guess to update value
		fmt.Scanln(&guess)

		// Check guess
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number")
			break

		} else if guess< target {
			fmt.Println("Too low! Try guessing a higher number")
		} else {
			fmt.Println("Too high! Try guessing a lower number")
		}
	}
}
