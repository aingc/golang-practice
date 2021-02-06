package main

import "fmt"

func main() {
	x := 5
	y := 10

	// common practice is to not use parenthesis, if you use it won't throw an error
	// but apparently common convention is to not use parenthesis?
	// If else
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than or equal to %d\n", y, x)
	}

	// else if
	color := "green"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("Color is not blue or red")
	}

	// Switch
	switch color {
		case "red":
			fmt.Println("color is red")
		case "blue":
			fmt.Println("color is blue")
		default:
			fmt.Println("color is not blue or red")
	}
}