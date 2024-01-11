package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Sam",
		Breed: "German Shephered",
	}
	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:          "Joe",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("The Animal Says ", a.Says(), " and has ", a.NumberOfLegs(), " legs")
}

func (d *Dog) Says() string {
	return "Bow"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (g *Gorilla) Says() string {
	return "Bow"
}

func (g *Gorilla) NumberOfLegs() int {
	return 4
}
