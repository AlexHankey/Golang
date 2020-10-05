package main

import "fmt"

func main()  {
	//s := make([]int, 3)
	//s[0] = 1
	//s[1] = 2
	//s[2] = 3
	//fmt.Println(s)
	//
	//fmt.Println(s[0])
	//
	//fmt.Println(len(s))
	//
	//s = append(s, 4, 5, 6, 7)
	//fmt.Println(s)
	//
	//fmt.Println(s[1:3])
	//fmt.Println(s[:3])
	//fmt.Println(s[0:])
	//
	//// concise slice definition
	//t := []int{100, 200, 300}
	//fmt.Println(t)
	//
	////x := s
	////x[0] = 500
	////fmt.Println(x)
	////fmt.Println(s)
	//
	//// copy prevents changing x and s
	//x := make([]int, len(s))
	//copy(x, s)
	//x[0] = 900
	//fmt.Println(x)
	//fmt.Println(s)
	//
	////2-d slice, similar to array although lengths of slice may vary
	//ss := make([][]int, 3)
	//for i := 0; i< 3; i++ {
	//	inner_len := i +1
	//	ss[i] = make([]int, inner_len)
	//	for j := 0; j < inner_len; j++ {
	//		ss[i][j] = i + j
	//	}
	//}
	//fmt.Println(ss)
	//
	//var myslice []int
	//printSlice(myslice)
	//s = append(myslice, 1)
	//printSlice(s)
	//
	//arr1 := make([]int, 10, 14)
	//arr1 = []int{10, 20, 30, 40}
	//arr1 = append(arr1, 100)
	//
	//arr1[2] = 99
	//printSlice(arr1)
	//
	////var vals [20]int
	////for i := 0; i < 5; i++ {
	////	vals[i] = i * i
	////}
	////subsetLen := 5
	////fmt.Println("The subset of our array has a length of:", subsetLen)
	////// add a new item to our array
	////vals[subsetLen] = 123
	////subsetLen++
	////fmt.Println("The subset of our array has a length of:", subsetLen)
	////////////////////////////
	//
	//var vals []int
	//for i := 0; i <5; i++ {
	//	vals = append(vals, i)
	//	fmt.Println("The length of our slice is:", len(vals))
	//	fmt.Println("The capacity of our slice is:", cap(vals))
	//}
	//// add a new item to our array
	//vals = append(vals, 123)
	//fmt.Println("The length of our slice is:", len(vals))
	//fmt.Println("The capacity of our slice is:", cap(vals))
	//// Accessing items is the same as an array
	//fmt.Println(vals[5])
	//fmt.Println(vals[2])

	vals := make([]int, 5)
	fmt.Println("Capacity was:", cap(vals))
	for i := 0; i < 5; i++ {
		vals = append(vals, i)
		fmt.Println("Capacity is now:", cap(vals))
	}

	fmt.Println(vals)

	fmt.Println(keys(map[string]struct{}{
		"dog": {},
		"cat": {},
	}))
}
func keys(m map[string]struct{}) []string {
	ret := make([]string, 0, len(m))
	for key := range m {
		ret = append(ret, key)
	}
	return ret
}
//func printSlice(s []int) {
//	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
//}
