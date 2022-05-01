package main

import "fmt"

func main() {
	age := 28
	name := "Sachin Rai"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("New Line \n")

	// Println
	fmt.Println("Hello Sachin!")
	fmt.Println("Bye Sachin!")

	fmt.Println("My name is", name, "and my age is", age)

	// Printf (formatted strings) %_ = format specifier [v,q,T,0.1f: decimal point]
	fmt.Printf("My name is %v and my age is %v \n", name, age)
	fmt.Printf("My name is %q and my age is %q \n", name, age)
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("you scored %0.1f points! \n", 255.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My name is %v and my age is %v \n", name, age)
	fmt.Println("The saved string is:", str)
}
