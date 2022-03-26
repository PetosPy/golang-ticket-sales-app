package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Michael's 2022"
	const ticketNumbers = 50
	var remainingTickets int = 50

	fmt.Printf("Hello and welcome to %v.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are available\n", ticketNumbers, remainingTickets)
	fmt.Println("Get your tickets here")

	var firstName string
	var lastName string
	var email string
	var ticketsBought int

	var bookings []string

	for remainingTickets > 0 && len(bookings) < 2 {

		remainingTickets = remainingTickets - ticketsBought

		fmt.Println("Please enter your first Name: ")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last Name: ")
		fmt.Scan(&lastName)

		fmt.Println("How many tickets would like to buy?")
		fmt.Scan(&ticketsBought)

		fmt.Println("Please enter your Email: ")
		fmt.Scan(&email)

		// Check if user input is valid.
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicket := ticketsBought > 0

		if isValidName && isValidEmail && isValidTicket {
			// Booking logic
			if ticketsBought <= remainingTickets {
				bookings = append(bookings, firstName+" "+lastName)

				fmt.Printf("Hello %v %v, thank you for booking %v tickets, recipt will be sent to %v\n", firstName, lastName, ticketsBought, email)
				fmt.Printf("We have %v tickets remaining for the %v conference\n", remainingTickets, conferenceName)

				firstNames := []string{}
				for _, booking := range bookings {
					var names = strings.Fields(booking)[0] // splits a string on the whitespaces
					firstNames = append(firstNames, names)
				}

				// fmt.Printf("This are all our bookings: %v\n", bookings)
				fmt.Printf("The first names of bookings are: %v\n", firstNames)

				if remainingTickets <= 0 {
					// end program
					fmt.Println("*** Our conference tickets are sold out! ***")
					break
				}

			} else {
				fmt.Printf("We only have %v tickets. Please try again!\n", remainingTickets)
			}

		} else {
			// Invalid data
			if !isValidName {
				fmt.Println("First name or last name you entered is too short!")
			}

			if !isValidEmail {
				fmt.Println("Please enter a valid Email address")
			}

			if !isValidTicket {
				fmt.Println("Please enter the right amount of tickets")
			}
		}
	}
}
