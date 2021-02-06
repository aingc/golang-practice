package main

import "fmt"

// Maps are key value pairs
func main() {
	// Define a map make(map[KEY]DATATYPE_OF_VALUE)
	// emails := make(map[string]string)

	// Assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["Sharon"] = "sharon@gmail.com"
	// emails["Mike"] = "sharon@gmail.com"

	// Can also assign key value pairs when defined
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	
}