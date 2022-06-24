package main

import (
	"fmt"
	"sync"
	"time"
)

// Package level constants
const conferenceTickets = 50
const conferenceName = "Go Conference"

// Package level variables
var remainingTickets uint = conferenceTickets

type UserData struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets uint
}

var bookings = make([]UserData, 0)
var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, email, numOFTickets := getUserInput()

		if !isValidUserInput(firstName, lastName, email, numOFTickets) {
			continue
		}

		bookTicket(numOFTickets, firstName, lastName, email)
		wg.Add(1)
		go sendTicket(numOFTickets, firstName, lastName, email)
	}

	fmt.Printf("The %v for this year has been sold out! Please try next year.\n", conferenceName)
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets for sale.\n", conferenceTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getUserInput() (string, string, string, uint) {
	var uFirstName string
	var uLastName string
	var uEmail string
	var uTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&uFirstName)
	fmt.Println("Enter your last name:")
	fmt.Scan(&uLastName)
	fmt.Println("Enter your email Id:")
	fmt.Scan(&uEmail)
	fmt.Println("Enter number of tickets you wish to book:")
	fmt.Scan(&uTickets)

	return uFirstName, uLastName, uEmail, uTickets
}

func bookTicket(uTickets uint, uFirstName string, uLastName string, uEmail string) {
	var userData = UserData{
		firstName:    uFirstName,
		lastName:     uLastName,
		email:        uEmail,
		numOfTickets: uTickets,
	}
	bookings = append(bookings, userData)
	remainingTickets -= uTickets
	if uTickets > 1 {
		fmt.Printf("Thank you %v %v. Your booking of %v tickets to %v is successful. Confirmation email will be sent to %v.\n", uFirstName, uLastName, uTickets, conferenceName, uEmail)
	} else {
		fmt.Printf("Thank you %v %v. Your booking of a ticket to %v is successful. Confirmation email will be sent to %v.\n", uFirstName, uLastName, conferenceName, uEmail)
	}

	fmt.Printf("The current bookings are: %v\n", getFirstNames())
	fmt.Printf("Number of tickets available for booking now: %v\n", remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(numOFTickets uint, firstName string, lastName string, email string) {
	time.Sleep(60 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v", numOFTickets, firstName, lastName)
	fmt.Println("##################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##################")
	wg.Done()
}
