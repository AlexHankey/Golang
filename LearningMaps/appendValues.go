//package main
//
//type person struct {
//	first 	string
//	last 	string
//	dob 	string
//}
//
//func main()  {
//
//	//x := []int{1, 2, 3, 4, 5}
//	//fmt.Println(x)
//	//
//	//p1 := person{
//	//	first: "Alex",
//	//	last:  "Hankey",
//	//	dob:   "22/07/1998",
//	//}
//	//fmt.Println(p1)
//	//
//	//j := map[int][]string {
//	//	1: "Alex",
//	//}
//	//fmt.Println(j)
//
//	//m := map[string][]string {
//	//	"cat": {"Brown", "Grey"},
//	//	"dog": {"black"},
//	//}
//	//add := append(m["dog"], "red")
//	//fmt.Println(add)
//	//
//	//m["fish"] = []string{"orange", "red"}
//	//fmt.Println(m)
//	//fmt.Println(m["fish"])
//
//
//}

package main

import (
	"fmt"
)
func main() {

	m := map[string][]string{

		"cat": {"Brown", "Grey"},

		"dog": {"black"},

	}
	m["dog"] = append(m["dog"], "red")

	fmt.Println(m)

	fmt.Printf("Map = %+v\n", m)

	fmt.Println("Why is the map unchanged?")

	mySlice1 := make([]int, 0)
	mySlice2 := []int{}
	fmt.Println("mySlice1", len(mySlice1))
	fmt.Println("mySlice2", len(mySlice2))
	
}