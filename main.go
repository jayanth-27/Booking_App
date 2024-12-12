package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to %v\n ", conferenceName)
	var userName string
	var userTickets int
	var userEmail string
	fmt.Println("Please enter your Name")
	fmt.Scan(&userName)

	fmt.Println("Please enter the number of tickets required")
	fmt.Scan(&userTickets)

	fmt.Println("Please enter your email")
	fmt.Scan(&userEmail)

	fmt.Printf("Thank you %v for booking %v tickets, you will receive the confirmation email to %v\n", userName, userTickets, userEmail)
	fmt.Printf("We have total capacity of %v out of which, %v are remaining", conferenceTickets, remainingTickets-userTickets)
}
