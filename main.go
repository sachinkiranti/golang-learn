package main

import "fmt"

func main() {

	// Array
	var ages [3]int = [3]int{20, 22, 28}
	var nums = [5]int{20, 22, 28, 32, 50}
	names := [4]string{
		"sachin",
		"kiranti",
		"uman",
		"kirant",
	}

	names[1] = "Kiranti Changed"

	fmt.Println(ages, nums, len(nums), names)

	// Slices (use array under the hood)
	var scores = []int{
		100,
		140,
		160,
	}

	scores[0] = 101
	scores = append(scores, 180)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "mula")
	fmt.Println(rangeOne, len(rangeOne))

}
