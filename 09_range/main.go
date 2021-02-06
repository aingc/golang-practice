package main

import "fmt"

// Range is used to loop through maps, arrays, slices, etc
func main() {
	ids := []int{33,76,54,23,11,2}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index, put an underscore in place of the i iterator for it to compile if
	// you are not using i within the for loop
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum", sum)

	// Range with map
	emails := map[string]string{"Bob":"bob@gmail.com", "Sharon":"sharon@gmail.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}

	for k := range emails {
		fmt.Println("Name: " + k)
	}

	for _, v := range emails {
		fmt.Println("Emails: " + v)
	}
}