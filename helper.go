package main

import (
	"strings"
)

func checkValidation(firstName string, lastName string, email string, userTicket int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remticket
	return isValidName, isValidEmail, isValidTicket
}
