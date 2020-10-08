package main

import (
	"fmt"
)

func main()  {
	country := map[string]int {"japan" : 1, "china" : 2, "canada" : 3}
	_, ok := country["africa"]
	if ok {
		fmt.Println("Key found value is: ")
	} else {
		fmt.Println("Key not found")
	}


	intMap := map[int]string {
		0: "zero",
		1: "one",
	}
	t := map[int]string{}
	if len(t) ==  {
		fmt.Println("\n Empty map")
	}
	if len(intMap) == nil {
		fmt.Println("\nEmpty map")
	} else {
		fmt.Println("\nNot empty map")
	}
	fmt.Println("\n ########### \n")

	//demonstrates both the nil check and the length
	//check that can be used for checking if a map is empty
	a := new(map[int64]string)
	if *a == nil {
		fmt.Println("empty")
	}
	fmt.Println(len(*a))
}