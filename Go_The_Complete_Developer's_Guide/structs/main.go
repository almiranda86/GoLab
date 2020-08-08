package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var p person

	p.firstName = "Andre"
	p.lastName = "Miranda"
	p.contactInfo.email = "myemail@email.com"
	p.contactInfo.zipCode = 000000

	pPointer := &p
	pPointer.updateName("Adao")
	p.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) print() {
	fmt.Printf("%+v", (*pointerToPerson))
}
