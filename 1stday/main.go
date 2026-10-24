package main

import "fmt"

func main() {

	// var name string;
	// name = "John Doe"
	// --or---

	// name :="John Doe"
    // --or---

	// var name= "John Doe"

	// --or---

	var(
		name string = "John Doe"
		age int = 30
	)

	fmt.Println("Hello, World!")
	fmt.Println(name)
	fmt.Println(age)
}
