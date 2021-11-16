package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const emptyString = ""
const spanishLanguage = "Spanish"


func Hello(name string, language string) string {
	if name == emptyString {
		name = "World"
	}

	if language == spanishLanguage {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
