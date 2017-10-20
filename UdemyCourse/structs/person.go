package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo contactInfo
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// left of on having to use pointers to deal with the appropriate copy within the cpu's memory
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
