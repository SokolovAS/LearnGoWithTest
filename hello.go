package main

import "fmt"

const englishHelloPrefix = "Hello, "


func Hello(s string) string {
	return englishHelloPrefix + s
}

func main() {
	fmt.Println("Hello world!")
}
