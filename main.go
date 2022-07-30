package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "The LORD'S conference"
	const totalTickets = 50
	var remTickets = 50
	var firstName string
	var lastName string
	var email string
	var amount int
	var bookingsList []string
	var error string
	//greetUsers(firstName, conferenceName)
	for remTickets > 0 {
		//fmt.Printf("Welcome to : %v \nTotal tickets %v of %v \n", conferenceName, remTickets, totalTickets
		fmt.Printf("Enter your  first name:")
		fmt.Scan(&firstName)
		fmt.Printf("Enter your  last name:")
		fmt.Scan(&lastName)
		fmt.Printf("Enter your  email:")
		fmt.Scan(&email)
		fmt.Printf("Enter your amount:")
		fmt.Scan(&amount)
		var email_string = "@"
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isEmailValid := strings.Contains(email, email_string)
		isValidTicketQ := amount > 0 && amount <= remTickets
		if isValidTicketQ && isEmailValid && isValidName {
			remTickets = remTickets - amount
			bookingsList = append(bookingsList, firstName+" "+lastName) // similar implementation in solidity, only with type reversed
			firstNames := []string{}
			for _, booking := range bookingsList {
				var names = strings.Fields(booking) // here we split the items of the array
				firstNames = append(firstNames, names[0])
			}
			noTicketRemaining := remTickets == 0
			if noTicketRemaining {
				fmt.Println("Our conference is booked out. come back next year")
				break

				fmt.Printf("Hello %v %v to %v, you bought %v tickets \n", firstName, lastName, conferenceName, amount)
				fmt.Printf("Remaining %v of %v\n", remTickets, totalTickets)
				fmt.Printf("The first names are: %v", firstNames)
				fmt.Println(bookingsList)
			} else if !isEmailValid {
				error = "Invalid email format\n, try adding @\n"
			} else if !isValidName {
				error = "Name must be atleast two(2) characters\n"
			} else if !isValidTicketQ {
				error = "Sorry you have either exceeded the quantity of tickets allowed\n"
			}
			fmt.Printf("Hello %v, %v", firstName, error)
			continue // continue keyword similar to that in python.

		}

	}
}
