//pass_fail reports whether a grade is passing or failing.
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {

	var status string
	var grade float64
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	if grade >= 60 {
		status = "pass"
	} else {
		status = "fail"
	}
	fmt.Println("A grade of", grade, "is", status)

}
