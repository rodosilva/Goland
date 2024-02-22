package main

import (
	"log"

	"github.com/rodosilva/Goland/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Some name"

	log.Println(myVar.TypeName)
}

// Remember on terminal you need to
// rodo@rodoDesktop:~/Goland$ go mod init github.com/rodosilva/Goland
// And that will create a go.mod
