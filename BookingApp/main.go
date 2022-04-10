package main

import "fmt"

func main() {
	var confName = "Go Conference"

	const confTickets = 50
	var confRemainingTickets uint = 30

	fmt.Println("Welcome to", confName, "booking application")
	fmt.Println("We have total of", confTickets, "tickets and", confRemainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var ticketNumber uint

	fmt.Printf("\n\nEnter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&ticketNumber)

	fmt.Printf("Thank you %v %v for booking %v tickets. A confirmation mail will be sent to %v\n", firstName, lastName, ticketNumber, email)
}
