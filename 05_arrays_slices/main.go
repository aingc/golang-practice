package main

import "fmt"

func main() {
	// Arrays: Have to be fixed length and name types
	// A slice is basically an array that doesn't have a fixed type

	// var fruitArr [2]string

	// Assign values
	// fruitArr[0] = "Apple"
	// fruitArr[1] = "Orange"

	// Declare and assign
	// fruitArr := [2]string{"Apple", "Orange"}

	// fmt.Println(fruitArr)
	// fmt.Println(fruitArr[1])

	fruitSlice := []string{"Apple", "Orange", "Grape", "Cherry"}

	fmt.Println(fruitSlice)
	// len() gives size of slice
	// if want range [startIdx:idxEndingUpTo]
	// so fmt.Println(fruitSlice[1:2]) would be just [Orange]
}