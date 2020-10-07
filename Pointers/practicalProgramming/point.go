package main

import "fmt"

//point to Memory address &
//Read through Memory address *

func main() {
	x := 15
	a := &x
	fmt.Println(a)
	fmt.Println(*a)

	*a = 5
	fmt.Println(x)
	*a = *a * *a // 5 * 5 = 25
	fmt.Println(x)
	fmt.Println(*a)
}