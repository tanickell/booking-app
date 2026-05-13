package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to the", conferenceName, "booking application!")
	// fmt.Println()
	fmt.Println("We have a total of", conferenceTickets, "tickets, with", remainingTickets, "still available.")
	fmt.Println("Get your tickets here to attend.")
}
