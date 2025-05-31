package main

import (
	"fmt"
	"sync"
	"time"
)

var conference = "Go conference"

const ticket int = 50

var remticket int = 50
var booking = make([]userData, 0)

type userData struct {
	firstname      string
	lastname       string
	email          string
	numberOfTicket int
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()
	for {

		firstName, lastName, email, userTicket := getUserInput()
		isValidName, isValidEmail, isValidTicket := checkValidation(firstName, lastName, email, userTicket)

		if isValidEmail && isValidName && isValidTicket {

			bookTicket(firstName, lastName, email, userTicket)
			wg.Add(1)
			go sendTicket(userTicket, firstName, lastName, email)
			firstname := getFirstName()
			fmt.Printf("The first name of bookings are: %v \n\n", firstname)

			if remticket == 0 {
				fmt.Printf("Out conference is booked out. Come back next year\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email entered doesn't have @ sign")
			}
			if !isValidTicket {
				fmt.Println("number of ticket you entered is invalid")
			}
		}

		wg.Wait()

	}
}

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conference)
	fmt.Printf("We have total of %v ticket and of out which %v is remaining\n", ticket, remticket)
}

func getUserInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var userTicket int

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)
	fmt.Println("Enter the number of tickets: ")
	fmt.Scan(&userTicket)
	return firstName, lastName, email, userTicket
}

func getFirstName() []string {
	firstname := []string{}
	for _, book := range booking {

		firstname = append(firstname, book.firstname)
	}
	return firstname
}

func bookTicket(firstName string, lastName string, email string, userTicket int) {

	var userData = userData{
		firstname:      firstName,
		lastname:       lastName,
		email:          email,
		numberOfTicket: userTicket,
	}

	booking = append(booking, userData)
	fmt.Printf("Thank you %v %v for booking %v ticket. You will recieve confirmation email at %v\n", firstName, lastName, userTicket, email)
	remticket -= userTicket
	fmt.Printf("%v ticket remaining for %v\n", remticket, conference)

	fmt.Printf("List of booking is %v\n", booking)
}

func sendTicket(userTicket int, firstName string, lastName string, email string) {

	time.Sleep(5 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("##############")
	fmt.Printf("Sending tickets:\n%v \nto email address %v\n", ticket, email)
	fmt.Println("##############")
	wg.Done()
}
