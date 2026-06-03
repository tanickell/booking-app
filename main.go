package main

import (
	"fmt"
	"strings"
)

func main() {

	conferenceName := "Go Conference" // var conferenceName string = "Go Conference" // var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{} // var bookings []string // var bookings = []string{}

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, with %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	for {

		var firstName string
		var lastName string
		var email string
		var userTickets uint

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

		// Input Validation
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("\nThank you %v %v for booking %v tickets!\n\nYou will receive a confirmation message at %v.\n", firstName, lastName, userTickets, email) // fmt.Println(userName)
			fmt.Printf("%v tickets remaining out of %v total for %v\n\n", remainingTickets, conferenceTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("The first names of our bookings are: %v\n\n", firstNames) // fmt.Printf("These are all our bookings: %v\n\n", bookings)

			if remainingTickets == 0 { // var noTicketsRemaining bool = remainingTickets == 0 // noTicketsRemaining := remainingTickets == 0
				// end the program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {

			// fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets.\n\n", remainingTickets, userTickets)
			fmt.Printf("Your input data is invalid. Please try again.\n\n")
		}
	}

	// print a goodbye message :D
	fmt.Printf("Have a nice day!\n\n")
}
