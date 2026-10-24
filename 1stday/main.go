package main

import "fmt"




func add(x int, y int) int {
	return x + y
}





func makecoffee(kind string,isSugar bool){
	fmt.Printf("Making a cup of %s coffee. Sugar: %t\n", kind,isSugar)
	

}







func main() {

	// var name string;
	// name = "John Doe"
	// --or---

	// name :="John Doe"
    // --or---

	// var name= "John Doe"

	// --or---


	// grouped variable declaration

	// var(
	// 	name string = "John Doe"
	// 	age int = 30
	// )

	// -----------------

	// multiple variable declaration

	// var x, y, z int = 1, 2, 3

	// --or---

	// var x, y, z int;
	// x=1
	// y=2
	// z=3


// ----------------const-------------------------------
	// const pi float64 = 3.14
	// const pi = 3.14



// --------------zero value-------------------------

// var age int  
// zero value is 0

// var name string
// zero value is empty string

// var isActive bool
// zero value is false

// var score float64
// zero value is 0.0

// -----------------------------------------


// var name string = "John Doe"
// age := 30





// -----------------------------------------



	fmt.Println("Hello, World!")
	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(x,y,z)
	// fmt.Println(pi)
	// fmt.Println(age)
	// fmt.Println(name)
	// fmt.Println(isActive)
	// fmt.Println(score)

	// ---------------------------------------

	// fmt.Printf("My name is %s and  I am %d years old", name, age)
	// result := fmt.Sprintf("My name is %s and  I am %d years old", name, age)
	// fmt.Println(result)


	// ------------------------------------




	makecoffee("Latte", true)



		// ------------------------------------



	result := add(3, 4)
	fmt.Println(result)
}




