package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, ticketsBought int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := ticketsBought > 0

	return isValidName, isValidEmail, isValidTicket
}
