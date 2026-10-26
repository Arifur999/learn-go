package main

import "fmt"

func makeCoffee(coffeeNo int) {
	fmt.Printf("Making cup %d of coffee\n", coffeeNo)
}

func main() {


	// for loop
	for i := 0; i <= 5; i++ {
		makeCoffee(i)
	}


		// while loop
	i := 0
	for i <= 5 {
		makeCoffee(i)
		i++
	}


	// infinite loop
	
	for {
		makeCoffee(2)
	}

	// break (breaks the loop)
	for i := 0; i <= 5; i++ {
		makeCoffee(i)
		if i == 3 {
			break
		}
	}


	// continue (skips the current iteration and continues with the next one)
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		makeCoffee(i)
	}
}