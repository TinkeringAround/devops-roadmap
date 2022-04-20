package main

import "fmt"

func main() {
	const totalDays uint = 90
	var daysCompleted uint

	fmt.Println("How many days have you completed of the 90 Days Challenge?")
	_, err := fmt.Scan(&daysCompleted)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("You have", totalDays-daysCompleted, "left to go!")
}
