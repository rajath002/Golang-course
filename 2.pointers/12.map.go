package main

import (
	"log"
	"sort"
)

func main() {
	myMap := make(map[string]string)
	var mySlice []string

	mySlice = append(mySlice, "Rajath")
	mySlice = append(mySlice, "Dravid")
	mySlice = append(mySlice, "Virat")

	myMap["dog"] = "Tommy"

	log.Println(myMap["dog"])

	numbers := []int{4, 2, 1, 8, 6, 9, 0}

	log.Println(mySlice)

	// Sorting
	sort.Ints(numbers[:])
	log.Println(numbers)

	// Reverse sorting
	sort.Sort(sort.Reverse(sort.IntSlice(numbers[:])))
	log.Println("Reverse sorting:", numbers)

}
