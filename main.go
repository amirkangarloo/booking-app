package main

import "fmt"

func main() {
	var availableTicketNumber uint = 50
	var firstName string
	var lastName string
	var email string

	const conferenceName = "Go conference (Booking app with go)"
	const totalTicketNumber uint = 50

	fmt.Printf("welcome to %v\n", conferenceName)
	fmt.Printf("\nTotal ticket is: %v tickets\nAvailable ticket is: %v tickets\n", totalTicketNumber, availableTicketNumber)

	fmt.Println("\nPlease enter your first name:")
	fmt.Scan(&firstName)
	fmt.Println("\nPlease enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("\nPlease enter your email:")
	fmt.Scan(&email)

	fmt.Printf("\nThanks %v! Book ticket successful ❤\nSee you ;)\n", firstName+" "+lastName)
}
