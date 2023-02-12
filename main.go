package main

import "fmt"

// struct type
type Person struct {
	firstName, lastName, email string
	tickets                    uint
}

// init global variables
var conferenceName string = "Golang"
var conferenceTickets uint = 30
var remainingTickets uint = 30
var bookings = make([]Person, 0)

func main() {
	// get user input
	firstName, lastName, email, tickets := getUserInput()
	// validate user input
	isValidFirstName, isValidLastName, isValidEmail, isValidTickets := validUserInput(firstName,
		lastName,
		email,
		tickets,
		remainingTickets)
	// book ticket
	// send ticket
	// change value remaining ticket
}

func getUserInput() (string, string, string, uint) {
	var firstName, lastName, email string
	var tickets uint
	fmt.Print("Input firstName:")
	fmt.Scan(&firstName)
	fmt.Print("Input lastName:")
	fmt.Scan(&lastName)
	fmt.Print("Input email:")
	fmt.Scan(&email)
	fmt.Print("Input tickets:")
	fmt.Scan(&tickets)
	return firstName, lastName, email, tickets
}
