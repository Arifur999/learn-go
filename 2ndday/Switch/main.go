package main

func main() {
	day := "friday"
	// if day == "friday" {
	// 	println("It's Friday! Time to relax and enjoy the weekend.")
	// }else if day == "saturday" {
	// 	println("It's Saturday! Time to have fun and do something you love.")
	// }


	// switch 


	switch  {
	case day == "monday":
		println("It's Monday! Time to start the week with energy and motivation.")
	case day == "tuesday":
		println("It's Tuesday! Keep going and stay focused on your goals.")
	case day == "wednesday":
		println("It's Wednesday! You're halfway through the week, keep pushing forward.")
	case day == "thursday":
		println("It's Thursday! The weekend is almost here, stay strong and finish the week strong.")
	case day == "friday":
		println("It's Friday! Time to relax and enjoy the weekend.")
	case day == "saturday":
		println("It's Saturday! Time to have fun and do something you love.")
	case day == "sunday":
		println("It's Sunday! Take some time to rest and recharge for the week ahead.")
	default:
		println("Invalid day of the week.")
	}








	// tagged switch

	switch day {
	case "monday":
		println("It's Monday! Time to start the week with energy and motivation.")
	case "tuesday":
		println("It's Tuesday! Keep going and stay focused on your goals.")
	case "wednesday":
		println("It's Wednesday! You're halfway through the week, keep pushing forward.")
	case "thursday":
		println("It's Thursday! The weekend is almost here, stay strong and finish the week strong.")
	case "friday":
		println("It's Friday! Time to relax and enjoy the weekend.")
	case "saturday":
		println("It's Saturday! Time to have fun and do something you love.")
	case "sunday":
		println("It's Sunday! Take some time to rest and recharge for the week ahead.")
	default:
		println("Invalid day of the week.")
	}
}