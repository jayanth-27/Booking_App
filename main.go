package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v\n ", conferenceName)
	fmt.Printf("We have total capacity of %v out of which, %v are remaining", conferenceTickets, remainingTickets)
}
