package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print(" Enter temp in Farenite: ")
	fahrenheit, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f deggress Celsius\n", celsius)

}
