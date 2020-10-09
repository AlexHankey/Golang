package main

import "fmt"

func main() {
	myInt := IntCounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 19; i ++ {
		fmt.Println(inc.Increment())
	}

}

type Incrementer interface {
	Increment() int
}
type Incrementer int

func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error)  {
	n, err := fmt.Println(string(data))
	return n, err
}