package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64
	// byte - alisa for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// Using var
	//var name = "John"
	// if you initialize a variable and you don't use it, there will be an error
	var age int = 37
	const isCool = true

	// Shorthand
	name:= "John"
	size:= 1.3
	// can also do
	// name, email := "Brad", "brad@gmail.com"

	fmt.Println(name, age, isCool)
	// format specifier
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", size)
}