package main

import (
	"log"

	"github.com/rajath002/mygopackage/helpers"
)

func main() {
	log.Println("hello")

	myVar := helpers.SomeType{
		TypeName:   "Some name",
		TypeNumber: 5,
	}

	var myVar2 helpers.SomeType
	myVar2.TypeName = "Google"
	myVar.TypeNumber = 6

	log.Println(myVar, myVar2)

}
