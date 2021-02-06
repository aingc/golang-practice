package main

import (
	"fmt"
	"strconv"
)

// Define a person structure
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int

	// can alternatively do
	// firstName, lastName, city, gender string
	// age int
}

// Greeting method (value receiver)
// func (identifier struct) METHOD_NAME() DATATYPE {}
func (person Person) greet() string {
	// if es6 or php, etc you'd use this keyword
	return "Hello, my name is " + person.firstName + " " + person.lastName + " and I am " + strconv.Itoa(person.age)

}

// hasBirthday method (pointer receiver) "we're going to change something"
func (person *Person) hasBirthday() {
	person.age++
}

// getMarried (pointer receiver)
func (person *Person) getMarried(spouseLastName string) {
	if person.gender == "m" {
		return
	} else {
		person.lastName = spouseLastName
	}
}

func main() {
	// Init a person using struct, this is more descriptive compared to the alternative below
	person1 := Person{firstName: "Samantha", lastName: "Smith", city: "Boston", gender: "f", age: 25}

	// Alternative initialization
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}

	fmt.Println(person1)

	// Print single field
	fmt.Println(person1.firstName)
	// Can alter fields
	person1.age++
	fmt.Println(person1)
	
	person1.hasBirthday()
	person1.getMarried("Williams")

	person2.getMarried("Thompson")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}