package main

import (
	"errors"
	"log"
)

func Devide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("Cannot devide by 0")
	}
	result = x / y

	return result, nil
}

func main() {

	result, err := Devide(5, 2)

	log.Println("Result: ", result, err)
}
