package main

import (
	"fmt"
	"strings"
)

// Package Variables -- Global
const conf_tickets = 50

var conf_name = "GO Conference"    // other way to write this if local --  var conf_name string = "GO conference" (---or---)  conf_name := "GO conference"
var remain_tickets uint = 50       // uint because the number cannot go in negative
var bookings = make([]UserData, 0) //this is a slice

// using struct data type instead of map
type UserData struct {
	firstName      string
	lastName       string
	email          string
	NumberofTicket uint
}

func main() {

	greetUser() // Calling function to Greet User the Package Variables will automatically called there

	for {
		// this is a infinte loop also you can give a conditional for by -- "for remain_tickets < 50{}"
		var firstName string
		var lastName string
		var email string
		var userTicket uint //uint

		// Takes User Input from the user for the Details
		fmt.Println("Enter your First Name : ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last Name :")
		fmt.Scan(&lastName)

		fmt.Println("Enter your Email :")
		fmt.Scan(&email)

		fmt.Println("Enter Number of Ticket you required : ")
		fmt.Scan(&userTicket)

		fmt.Println(" ")

		isValidName, isValidEmail, isValidCount := ValidUserInput(firstName, lastName, email, userTicket, remain_tickets)

		if isValidName && isValidEmail && isValidCount {

			remain_tickets = remain_tickets - userTicket

			// calling struct that we created
			var userData = UserData{
				firstName:      firstName,
				lastName:       lastName,
				email:          email,
				NumberofTicket: userTicket,
			}

			bookings = append(bookings, userData)
			fmt.Printf("The full details of the User : %v \n", bookings)

			fmt.Printf("Thank You %v %v for booking %v Tickets... You will receive mail to %v \n", firstName, lastName, userTicket, email)
			fmt.Printf("The remaining Tickets available -- %v ... \n", remain_tickets)
			fmt.Println(" ")

			// add go in front of the function sendTicket to make it as new thread and remove the comments inside that function for 10 sec delay
			sendTicket(firstName, lastName, userTicket, email) // go (goroutine) -  a new thread

			// Calling function to print the First Name
			loopFirstNames := printFirstName()
			fmt.Printf("The First Names of the bookings for the Conference are : %v \n", loopFirstNames)
			fmt.Println(" ")

			if remain_tickets == 0 {
				fmt.Println("The tickets for our conference got over ")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("The name you entered is too short...")
			}
			if !isValidEmail {
				fmt.Println("The entered email is invalid and does not contain @ in it...")
			}
			if !isValidCount {
				fmt.Println("The number of Tickets you entered is invalid ")
			}
		}
	}
}

func greetUser() {
	fmt.Println("Welcome to our", conf_name, " booking application")
	fmt.Println("we have total of", conf_tickets, "and the remaing tickets are ", remain_tickets)
	//fmt.Printf("Welcome to our %v booking application \n", conf_name)
	fmt.Println("Get your tickets here.....")
	fmt.Println(" ")
}

func printFirstName() []string {
	var loopFirstNames = []string{} // other way is  -- "var firstNames []string"
	for _, i := range bookings {
		loopFirstNames = append(loopFirstNames, i.firstName)
	}
	return loopFirstNames
}

func ValidUserInput(firstName string, lastName string, email string, userTicket uint, remain_tickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail = strings.Contains(email, "@")
	isValidCount := userTicket > 0 && userTicket < remain_tickets

	return isValidName, isValidEmail, isValidCount
}

func sendTicket(firstName string, lastName string, userTicket uint, email string) {

	// time.Sleep(10 * time.Second) // Adding 10 sec delay to work on Concurrency
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName)
	fmt.Println("###############")
	fmt.Printf("Sending Ticket :\n %v \nto email %v\n", ticket, email)
	fmt.Println("###############")
}
