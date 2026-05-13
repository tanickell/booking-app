package main

import "fmt"

func main() {
	conferenceName := "Go Conference" // var conferenceName string = "Go Conference" // var conferenceName = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	fmt.Printf("Welcome to the %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, with %v still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int
	// ask the user for their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&userName)

	// fmt.Println(remainingTickets)
	// fmt.Println(&remainingTickets)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets) // fmt.Println(userName)
}
