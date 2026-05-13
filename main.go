package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // var conferenceName string = "Go Conference" // var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, with %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask the user for their first name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	// ask the user for their last name
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	// ask the user for their email address
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	// ask the user for number of tickets desired
	fmt.Println("Enter number of tickets desired: ")
	fmt.Scan(&userTickets)

	fmt.Printf("\nThank you %v %v for booking %v tickets!\n\nYou will receive a confirmation message at %v.\nHave a nice day!\n", firstName, lastName, userTickets, email) // fmt.Println(userName)
}
