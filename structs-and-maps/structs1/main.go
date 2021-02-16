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
	greg := person{
		firstName: "Greg",
		lastName:  "Davies",
		contactInfo: contactInfo{
			email:   "taskmaster@gmail.com",
			zipCode: 99999,
		},
	}

	greg.updateName("Gregory")
	greg.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println(&p)
}
