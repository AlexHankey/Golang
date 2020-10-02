package main

import "fmt"

func main() {
	//specify key/value pair
	m := make(map[string]int)

	m["a"] = 0
	m["b"] = 1

	fmt.Println(m)

	// print value of specific key
	fmt.Println("Value of key a in map:", m["a"])

	// remove key pair from map. which can be done with delete key-word
	delete(m,"a")
	fmt.Println(m)

	// another way to initialize a map (if you alreayd know the values/keys ahead of time
	m2 := map[string]int {"Key1": 1, "Key2": 2}
	fmt.Println(m2)

	// the value and whether the value is present
	val, isValPresent := m["b"]
	fmt.Println(val)
	fmt.Println(isValPresent)

	// check if something is not present
	_, isValPresent2 := m["a"]
	fmt.Println(isValPresent2)



}