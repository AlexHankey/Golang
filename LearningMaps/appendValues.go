package main

import "fmt"

type person struct {
	first 	string
	last 	string
	dob 	string
}

func main()  {

	//x := []int{1, 2, 3, 4, 5}
	//fmt.Println(x)
	//
	//p1 := person{
	//	first: "Alex",
	//	last:  "Hankey",
	//	dob:   "22/07/1998",
	//}
	//fmt.Println(p1)
	//
	//j := map[int][]string {
	//	1: "Alex",
	//}
	//fmt.Println(j)

	m := map[string][]string {
		"cat": {"Brown", "Grey"},
		"dog": {"black"},
	}
	res := append(m["dog"], "red")
	fmt.Println(res)

	m["fish"] = []string{"orange", "red"}
	fmt.Println(m["fish"])

}