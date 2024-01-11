package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	HasDog    bool   `json:"hasDog"`
}

func main() {
	myJson := `
	[
		{
			"firstName": "Clark",
			"lastName": "Kent",
			"hasDog": true
		},
		{
			"firstName": "Bruce",
			"lastName": "Wayne",
			"hasDog": false
		}
	]
	`
	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error while unmarshalling JSON ", err)
	} else {
		log.Println("Unmarshalled: ", unmarshalled)
	}

	// Write JSON from a struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Wally"
	m2.LastName = "West"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "")
	if err != nil {
		log.Println("Error marshalling ", err)
	}

	fmt.Println(string(newJson))
}
