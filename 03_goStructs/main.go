package main

import "fmt"

type contactInfo struct {
	email string
	telephone string
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	john := person{
		firstName: "John",
		lastName: "Smith",
		contactInfo: contactInfo{
			email: "john@email.com",
			telephone: "0123",
		},
	}

	john.updateName("Jim")
	john.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}