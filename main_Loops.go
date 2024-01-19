package main

import "log"

func main() {

	for i := 0; i <= 5; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	animals2 := make(map[string]string)
	animals2["dog"] = "Fido"
	animals2["cat"] = "Fluffy"

	for animalType, animal := range animals2 {
		log.Println(animalType, animal)
	}

	var firstLine = "Once upon"

	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@mail.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@mail.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@mail.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@mail.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}
