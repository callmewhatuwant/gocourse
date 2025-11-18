package main

import "fmt"

func main() {

	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "test@test.de",
	}

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
	}

	p1 := Person{
		firstName: "Jane",
		age:       25,
	}

	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p, p1)

	fmt.Println(user)

	fmt.Println(p.fullName())

}

type Person struct {
	firstName string
	lastName  string
	age       int
}

func (p Person) fullName() string{
	return p.firstName + " " + p.lastName
	
}
