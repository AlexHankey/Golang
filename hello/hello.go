package main

import "fmt"

const spanish = "Spanish"
const french = "french"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = " Bonjor, "

func Hello(name string, language string) string {
	if name == "" {
		name = "Alex"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frenchHelloPrefix + name
	}


	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}