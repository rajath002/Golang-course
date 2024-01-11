package main

import (
	"log"
	"time"
)

type User struct {
	FirstName   string
	Lastname    string
	PhoneNumber int
	Age         int
	BirthDate   time.Time
}

// You can ignore the warning
func main() {
	user := User{
		FirstName: "Rajath",
		Lastname:  "Kumar",
		BirthDate: time.Now(),
	}
	log.Println("Output", user.FirstName, user.Lastname, user.BirthDate)
}
