package main

import "fmt"

type contact struct {
	email string
	phone int64
}

type person struct {
	fName string
	lName string
	contact
}

func main() {
	jacob := person{
		fName: "Parser",
		lName: "Jacob",
		contact: contact{
			email: "parserjacob@yahoo.com",
			phone: 9876543210,
		},
	}

	jacob.print()
	jacob.changeName("Jacob", "Parser")
	jacob.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) changeName(fName string, lName string) {
	(*p).fName = fName
	(*p).lName = lName
}
