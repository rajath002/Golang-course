package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "This is the Home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, "This is the About page & 2 + 2 is : %d", sum)
}

func AddValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x := float32(100.0)
	y := float32(0.0)
	f, err := DevideValues(x, y)
	if err != nil {
		fmt.Fprintf(w, "Cannot devide by 0")
		return
	}
	fmt.Fprintf(w, "%f devided by %f is %f", x, y, f)
}

func DevideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}

func main() {

	http.HandleFunc("/", Home)

	http.HandleFunc("/about", About)

	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
