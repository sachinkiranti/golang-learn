package main

import (
	"fmt"
)

func main() {

	x := 0

	for x < 5 {
		fmt.Println("Value of x is :", x)
		x++
	}

	// Traditional For loop
	for i := 0; i < 10; i++ {
		fmt.Println("Value of i is :", i)
	}

	names := []string{"sachin", "rai", "kiranti", "uman", "sweata", "jharna"}
	for name := 0; name < len(names); name++ {
		fmt.Println(names[name])
	}

	for index, value := range names {
		fmt.Printf("The value at index %v is %v\n", index, value)
	}

	for _, value := range names {
		fmt.Printf("The value is %v\n", value)
	}

	fmt.Println(names)

}
