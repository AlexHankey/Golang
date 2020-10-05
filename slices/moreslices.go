package main

import (
	"fmt"
)

func main()  {

	zero := []int{}
	fmt.Println(zero)
	zero = append(zero, 100)
	fmt.Printf("Look into the slice: %v \n", zero)
	fmt.Printf("appended to a zero length slice details: len:%d cap: %d %d \n", len(zero), cap(zero), zero)


	slice := make([]int, 0, 6)
	anotherslice := make([]int, 0, 3)

	slice = []int{10, 20, 30}
	anotherslice = []int{40, 50, 60}
	// Trailing dots ... = Variatic function
	// They split up the slice and append
	// example:
	// newSlice := make([]int, 0, 10)
	// newSlice = append(slice, anotherslice[0], anotherslice[2], anotherslice[2])
	slice = append(slice, anotherslice...)

	fmt.Printf("combined slices: len: %d / cap: %d", len(slice), cap(slice), slice)
}