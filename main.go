package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // var conferenceName string = "Go Conference" // var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// remainingTickets = -1 // won't work bc remainingTickets is uint type (see above)

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, with %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int
	// ask the user for their name

	userName = "Tom"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets) // fmt.Println(userName)
}
