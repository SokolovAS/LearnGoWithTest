package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const emptyString = ""
const spanishLanguage = "Spanish"
const frenchLanguage = "French"

func Hello(name string, language string) string {
	if name == emptyString {
		name = "World"
	}

	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanishLanguage:
		prefix = spanishHelloPrefix
	case frenchLanguage:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("", frenchLanguage))
}
