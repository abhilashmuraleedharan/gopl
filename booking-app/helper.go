package main

import (
	"fmt"
	"strings"
)

func isValidUserInput(firstName string, lastName string, email string, uTickets uint) bool {
	result := true
	if len(firstName) < 2 {
		fmt.Println("First name is too short. Should have at least two characters. Try again.")
		result = false
	}

	if len(lastName) < 1 {
		fmt.Println("Last name cannot be empty, Try again.")
		result = false
	}

	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	if !isValidEmail {
		fmt.Println("Invalid email. Email should contain a \"@\" and a \".\"")
		result = false
	}

	if uTickets > remainingTickets {
		if remainingTickets > 1 {
			fmt.Printf("You cannot book %v tickets since we have only %v tickets left. Try again\n", uTickets, remainingTickets)
		} else {
			fmt.Printf("You cannot book %v tickets since we have only 1 ticket left. Try again\n", uTickets)
		}
		result = false
	}

	return result
}
