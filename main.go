package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var conferenceTickets int = 50
var remainingTickets uint = 50
var conferenceName = "GoLang Conference"
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var secondName string
	var email string
	var userTickets uint

	fmt.Println("Please enter your firstName")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your secondName")
	fmt.Scan(&secondName)

	fmt.Println("Please enter your email adress")
	fmt.Scan(&email)

	fmt.Println("Please enter your userTickets")
	fmt.Scan(&userTickets)

	return firstName, secondName, email, userTickets
}

func main() {

	cnt := 0
	for {
		greetUsers()

		firstName, secondName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, secondName, email, userTickets)

		wg.Add(1)
		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, secondName, email)

			go sendTicket(userTickets, firstName, secondName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
		cnt += 1

		if cnt == 3 {
			break
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(userTickets uint, firstName string, secondName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        secondName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, secondName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
