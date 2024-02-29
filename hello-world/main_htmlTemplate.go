package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// const can't be changed
const portNumber = ":8080"

// In order for a function to respond for a web browser
// It needs to have 2 parameters
// Response parameter and a request

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil) // If there is an error, I don't care. That means _

}

// To create a Go module. This app uses Go
// Import third party packages
// This app uses all the standard library and also uses external dependencies
// $ go mod init myapp
