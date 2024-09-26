package main

import "fmt"

func main() {
	Greeting()
	advanceGreeting("Aman")
	advanceGreeting("Azure")
}

func Greeting() {
	fmt.Println("Hey There !!")
}

func advanceGreeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
