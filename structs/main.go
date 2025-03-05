package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time

}

func main() {
	user := User {
		FirstName: "Peter",
		LastName: "Teszary",
	}

	log.Println(user.FirstName)
	log.Println(user.LastName)
}