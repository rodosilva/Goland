package main

import (
	"errors"
	"fmt"
	"net/http"
)

// const can't be changed
const portNumber = ":8080"

// In order for a function to respond for a web browser
// It needs to have 2 parameters
// Response parameter and a request

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum))
}

// addValues adds two integers and return the sum
func addValues(x, y int) int {

	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "Can not divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0, f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Can not divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil

}

// main is the main application function
func main() {
	// Commenting this first lesson
	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		n, err := fmt.Fprintf(w, "Hello world!")
	//		if err != nil { //If there is an error. Nil means no error
	//			fmt.Println(err)
	//		}
	//		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n)) //%d is going to be replaced by n
	//	})

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) // If there is an error, I don't care. That means _

}

// To create a Go module. This app uses Go
// Import third party packages
// This app uses all the standard library and also uses external dependencies
// $ go mod init myapp
