package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	for remainingTickets != 0 {
		booking := []string{}
		//var booking[50]string for array of size 50
		//Slice
		fmt.Printf("Welcome to %v\n ", conferenceName)
		var userName string
		var userTickets int
		var userEmail string
		fmt.Println("Please enter your Name")
		fmt.Scan(&userName)
		fmt.Println("Please enter the number of tickets required")
		fmt.Scan(&userTickets)
		if userTickets > remainingTickets {
			fmt.Printf("Only %v tickets are available, unable to book %v tickets\n", remainingTickets, userTickets)
			continue
		}
		fmt.Println("Please enter your email")
		fmt.Scan(&userEmail)

		isValidName := len(userName) > 3
		isValidTicket := userTickets > 0
		isValidEmail := strings.Contains(userEmail, "@")
		if isValidName && isValidEmail && isValidTicket {
			booking = append(booking, userName)
			remainingTickets -= userTickets
			fmt.Printf("Thank you %v for booking %v tickets, you will receive the confirmation email to %v\n", userName, userTickets, userEmail)
			if remainingTickets == 0 {
				fmt.Printf("Tickets are sold out")
				break
			}
			fmt.Printf("We have total capacity of %v out of which, %v are remaining\n", conferenceTickets, remainingTickets)
			fmt.Printf("Users are %v\n", booking)
		} else {
			if !isValidName {
				println("Invalid user name")
			}
			if !isValidEmail {
				println("Invalid email")
			}
			if !isValidTicket {
				println("Invalid number of tickets")
			}
		}
	}
}
