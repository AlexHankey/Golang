// Ranging over a map of employees
package main

import "fmt"

func main() {
	var employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}
	for key, value := range employee {
		fmt.Println("Key:", key, "=>", "Element:", value)
	}
}