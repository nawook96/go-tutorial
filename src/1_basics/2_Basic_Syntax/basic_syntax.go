package main

import "fmt"

func main() {
	// Bracket
	i := 10
	if i > 5 {
		fmt.Println("above 5")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// Semicolon
	fmt.Println("Hello, world!")
	fmt.Println("Hello,")
	fmt.Println("world!")

	// Comment out
	// fmt.Println("Hello, world!")
	/*
		fmt.Println("Hello,")
		fmt.Println("world!")
	*/
}
