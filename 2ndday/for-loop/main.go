package main

import "fmt"

func makeCoffee(coffeeNo int) {
	fmt.Printf("Making cup %d of coffee\n", coffeeNo)
}

func main() {

	for i := 0; i <= 5; i++ {
		makeCoffee(i)
	}
}