package main

import "fmt"

func FizzBuzz(number int) {
	var num_series []int
	for i := 0; i < number+1; i++ {
		num_series = append(num_series, i)
		// check if i is div by 3 and 5
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(num_series[i])
		}

	}

}

func main() {
	FizzBuzz(15)
}
