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

	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Adress{
			city:    "London",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "2639295",
			cell: "327698265",
		},
	}

	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	p3 := Person{
		firstName: "Jane",
		age:       25,
	}

	//p2.address.city = "New York"
	//p2.address.country = "USA"

	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)
	fmt.Println(p1, p2)
	fmt.Println(p1.address.country)
	fmt.Println(p2.address.country)
	fmt.Println(p1.cell)
	fmt.Println(p1.home)

	fmt.Println(user)

	fmt.Println(p1.fullName())

	// p2.incrementAgeByOne()
	// fmt.Println(p2.age)

	fmt.Println("Are p1 and p2 equal:", p1 == p2)
	fmt.Println("Are p2 and p3 equal:", p2 == p3)

}

type Person struct {
	firstName string
	lastName  string
	age       int
	address   Adress
	PhoneHomeCell
}

type PhoneHomeCell struct {
	home string
	cell string
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
