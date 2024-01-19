package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	//Maps

	var myString string
	var myInt int

	myString = "Hi"
	myInt = 11

	mySecondString := "another string"

	log.Println(myString, mySecondString, myInt)

	// var myOtherMap map[string]string //Another way to declare
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "Fido"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]int)

	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"], myMap2["Second"])

	myMap3 := make(map[string]User)

	me := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap3["me"] = me

	log.Println(myMap3["me"].FirstName)

	//Slices

	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[0:2])
	log.Println(numbers[6:9])

	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)

}
