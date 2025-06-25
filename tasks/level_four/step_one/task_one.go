// "Арсений как человек"

package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) PrettyPrint() {
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s", p.name, p.age, p.address)
}
