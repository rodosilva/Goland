package main

import (
	"fmt"
	"myapp/pkg/handlers"
	"net/http"
)

// const can't be changed
const portNumber = ":8080"

// In order for a function to respond for a web browser
// It needs to have 2 parameters
// Response parameter and a request

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) // If there is an error, I don't care. That means _

}

// To create a Go module. This app uses Go
// Import third party packages
// This app uses all the standard library and also uses external dependencies
// $ go mod init myapp
