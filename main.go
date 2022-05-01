package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greeting := "Hello there friends!"

	// Using strings lib
	// strings does not modify
	// return bool
	fmt.Println(strings.Contains(greeting, "Hello there"))

	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))

	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Title(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// Using sort lib
	// sort does modify
	ages := []int{2, 5, 8, 1, 40, 30}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"sachin", "rai", "kiranti", "uman", "sweata", "jharna"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "sweata"))

}
