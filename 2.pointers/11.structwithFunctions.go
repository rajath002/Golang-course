package main

import (
	"log"
	"time"
)

type myStruct struct {
	Name        string
	Dateofbirth time.Time
}

func (m *myStruct) GetAge() int {
	currentDate := time.Now()
	age := currentDate.Year() - m.Dateofbirth.Year()
	// Adjust the age if the birthday hasn't occurred yet this year
	if currentDate.YearDay() < m.Dateofbirth.YearDay() {
		age--
	}
	return age
}

func main() {
	user := myStruct{
		Name:        "John Doe",
		Dateofbirth: time.Date(1995, time.March, 12, 0, 0, 0, 0, time.UTC),
	}

	log.Println("User info : ", user.Name, ", Date of birth : ", user.GetAge())
}
