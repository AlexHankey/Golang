package main

import (
	"fmt"
	"reflect"
)

func main() {
	//s := []string {
	//	"apple",
	//	"grape",
	//	"orange",
	//}
	//fmt.Println(s)
	//
	//b := contains(s, "apple")
	//fmt.Println(b)
	//
	//b = contains(s, "banana")
	//fmt.Println(b)

	items := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(SliceExists(items, 5))
	fmt.Println(SliceExists(items, 10))
}
func SliceExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		panic("SliceExists() given a non-slice type")
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}






func contains(f []string, n string) bool  {
	for _, ok := range f {
		if ok == n {
			return true
		}
	}
	return false
}