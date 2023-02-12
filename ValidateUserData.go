package main

import (
	"regexp"
)

func validUserInput(firstName string, lastName string, email string, tickets uint, remainTickets uint) (bool, bool, bool, bool) {
	return isValidName(firstName), isValidName(lastName), isValidEmail(email), isValidTickets(tickets, remainTickets)
}

func isValidName(name string) bool {
	match, _ := regexp.MatchString("[^A-Za-z+]", name)
	return match
}

func isValidEmail(email string) bool {
	match, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email)
	return match
}

func isValidTickets(tickets uint, remainTickets uint) bool {
	return tickets > 0 && tickets <= remainTickets
}
