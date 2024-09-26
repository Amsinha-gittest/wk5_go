package main

import "fmt"

func main() {
	Greeting()
	advanceGreeting("Aman")
	advanceGreeting("Sinha")
	fmt.Println(add(1, 2))
	add(1, 2)
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %d, REminder: %d\n", q, r)
	q, r = 20, 3
	fmt.Printf("Quotient: %d, REminder: %d\n", q, r)
}
