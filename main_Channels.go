package main

import (
	"log"

	"github.com/rodosilva/Goland/helpers"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber

}

func main() {

	intChan := make(chan int)
	defer close(intChan) // Whatever comes after "defer", execute that as soon as the current function is done

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)

}
