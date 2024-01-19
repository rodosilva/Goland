package main

import (
	"log"
	"time"
)

var s = "seven"

// var firstName string
// var lastName string
// var phoneNumber string
// var age int
// var birthDate time.Time
// Type struct can replace individual variables

type User struct {
	FirstName   string //Capital letter to be visible from outside
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {

	var s2 = "six"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")

	user := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.BirthDate)

}

func saySomething(s string) (string, string) {
	log.Println("s from the saySomething func is", s)
	return s, "World"
}
