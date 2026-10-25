package main

import "fmt"

// func makeCoffee(kind string) string {

// 	coffee := fmt.Sprintf("%s", kind)
// 	return coffee
// }

func makeCoffee(kind string) (coffee string, price int) {

	coffee = fmt.Sprintf("%s", kind)
	price = 5 // Example price, replace with actual pricing logic
	return coffee, price
}

func main() {



// anonymous function
	 coffeeShop := func() {
	 	fmt.Printf("Hello, World!")
	}
	coffeeShop()



// immediately invoked function
	 func() {
	 	fmt.Printf("Hello, World!")
	 }()





	 
	fmt.Println("Hello, World!")
	// fmt.Println(makeCoffee("arabica"))
	// fmt.Println(makeCoffee("robusta"))
	fmt.Println(makeCoffee("arabica"))
}
