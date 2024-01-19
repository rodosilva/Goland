package main

import "log"

func main() {

	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)
	changeUsingPointer(&myString) // & means pass this reference to this pointer to this function
	log.Println("afert fun call myString is set to", myString)
}

func changeUsingPointer(s *string) { //* is a pointer to a string. Pointer to a location in memory
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue //Pointer to s. * Refer to an actuall pointer
}
