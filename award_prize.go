package main

import (
	"fmt"
	"math/rand"
	"time"
)

func awardPrize() {
	doornumber := rand.Intn(3) + 1
	if doornumber == 1 {
		fmt.Println("You win a cruise!")
	} else if doornumber == 2 {
		fmt.Println("You win a car!")
	} else if doornumber == 3 {
		fmt.Println("You win a goat!")
	} else {
		panic("invalid door number")
	}
}

func main() {
	rand.Seed(time.Now().Unix())
	awardPrize()
}
