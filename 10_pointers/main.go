package main

import "fmt"

func main() {
	a := 5
	b := &a // assigning b to a pointer of a

	fmt.Println(a, b)
	fmt.Printf("%T %T\n", a, b) // int and *int
	
	// Like in C++ to read the address of the integer pointer, you use "address of" &,
	// use * for getting value
	// or do *& to get address of a from b pointer, aka the address that b is pointing to
	fmt.Println(*&b)

	// *b is *&a
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)

	// if you choose to pass addr instead of data itself, it can increase performance
	// roughly everything in go is passed by value
}