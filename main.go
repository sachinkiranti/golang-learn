package main

import "fmt"

var someName = "Hello"

func main() {
	// String
	var nameOne string = "sachin"
	// fmt.Println(nameOne)
	var nameTwo = 5
	var nameThree string

	nameOne = "Changed"
	nameThree = "Declared"

	nameFour := "Kumar" // Short hand version and won't work outside the method

	fmt.Println(nameOne, nameTwo, nameThree, nameFour, someName)

	// Integer
	var intOne int = 5
	intTwo := 10
	fmt.Println(intOne, intTwo)

	// Bits & Memory
	var numOne int8 = 20
	var numTwo int8 = -128
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	// Float
	var scoreOne float32 = -1.1221
	var scoreTwo float64 = 213123213123123123123123123123123211.12
	scoreThree := 12121211
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
