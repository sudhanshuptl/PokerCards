package main

import "fmt"

type person struct {
	firstName string
	lastname  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipcode int
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName(firstName string) {
	p.firstName = firstName
}
