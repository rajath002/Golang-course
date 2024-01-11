package main

import (
	"log"
	"math/rand"
)

func RandomNumber(n int) int {
	// rand.Seed(time.Now().UnixNano()) = deprecated
	rand.NewSource(86)
	value := rand.Intn(n)
	return value
}

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := RandomNumber(10)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
