package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const portuguesHelloPrefix = "Ola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case portuguese:
		prefix = portuguesHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
