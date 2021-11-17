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
	if language == spanishLanguage {
		return spanishHelloPrefix + name
	}
	if language == frenchLanguage {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", frenchLanguage))
}
