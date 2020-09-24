package main

import "fmt"

const spanish = "Spanish"
const french = "french"
const japanese = "japanese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = " Bonjor, "
const japaneseHelloPrefix = " Konichiwa, "

func Hello(name string, language string) string {
	if name == "" {
		name = "User"
	}

	prefix := englishHelloPrefix

	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case japanese:
		prefix = japaneseHelloPrefix

	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}