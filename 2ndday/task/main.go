package main

import "fmt"

// menu:
// 1)calculator grade
// 2)check pass/fail status
// 3)exit
// Enter your choice: 3
// exiting the program...goodbye!

func main() {

	// age := 20
	// nid := true

	// if age >= 18 && nid == true {
	// 	println("You are eligible to vote")
	// } else {
	// 	println("You are not eligible to vote")
	// }

	dispalyMenu := func() {
		fmt.Println("welcome to the grade calculator")
		fmt.Println("1) Calculate Grade")
		fmt.Println("2) Check Pass/Fail Status")
		fmt.Println("3) Exit")
		fmt.Print("Enter your choice: ")

	}

	var choice int

	running := true
	for running {
		dispalyMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var score int
			fmt.Print("Enter your score: ")
			fmt.Scan(&score)
			grade := calculatorGrade(score)
			fmt.Printf("Your grade is: %s\n", grade)
		case 2:
			var score int
			fmt.Print("Enter your score: ")
			fmt.Scan(&score)
			status := checkPassFail(score)
			fmt.Printf("Your status is: %s\n", status)
		case 3:
			fmt.Println("Exiting the program...goodbye!")
			running = false
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func calculatorGrade(score int) string {
	if score >= 90 || score == 100 {
		return "A"
	} else if score >= 80 && score <= 89 {
		return "B"
	} else if score >= 70 && score <= 79 {
		return "C"
	} else if score >= 60 && score <= 69 {
		return "D"
	} else {
		return "F"
	}
}

func checkPassFail(score int) string {
	if score >= 60 {
		return "Pass"
	}
	return "Fail"
}
