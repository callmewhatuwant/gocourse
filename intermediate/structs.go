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
		address: Adress{
			city: "London",
			country: "U.K.",
		},
	}

	p1 := Person{
		firstName: "Jane",
		age:       25,
	}

	p1.address.city = "New York"
	p1.address.country = "USA"


	fmt.Println(p.firstName)
	fmt.Println(p1.firstName)
	fmt.Println(p, p1)
	fmt.Println(p.address.country)
	fmt.Println(p1.address.country)

	fmt.Println(user)

	fmt.Println(p.fullName())

	p1.incrementAgeByOne()
	fmt.Println(p1.age)

}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Adress
}

type Adress struct {
	city    string
	country string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() {
	p.age++

}
