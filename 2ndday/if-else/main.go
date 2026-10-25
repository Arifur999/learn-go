package main

import "fmt"

func main() {
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult. You are", age, "years old.")
	} else if age >= 13 {
		fmt.Println("You are a teenager. You are", age, "years old.")
	}else if age >= 0 {
		fmt.Println("You are a child. You are", age, "years old.")
	}else {
		fmt.Println("You are not an adult. You are", age, "years old.")
	}
}