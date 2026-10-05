package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactinfo
}

type contactinfo struct {
	email string
	zip   int
}

func (p person) print() {
	fmt.Println(p.lastName + ", " + p.firstName)
	fmt.Printf("%+v", p.contact)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	fmt.Println("Hello, Golang Structs!")
	var p person
	xl := person{
		firstName: "Michael",
		lastName:  "Jackson",
	}
	fmt.Println(xl)
	xl.print()
	p.firstName = "Q"
	p.lastName = "H"
	p.contact.email = "qh@example.com"
	p.contact.zip = 12345
	p.print()

	p.updateName("John")
	p.print()
}
