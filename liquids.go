package main

import "fmt"

type Gallons float64

func (g Gallons) String() string {
	return fmt.Sprintf("%0.2f gal", g)
}

type Liters float64

func (l Liters) String() string {
	return fmt.Sprintf("%0.2f L", l)
}

type Mililiters float64

func (ml Mililiters) String() string {
	return fmt.Sprintf("%0.2f mL", ml)
}

func main() {
	fmt.Println(Gallons(12.09248342))
	fmt.Println(Liters(12.09248342))
	fmt.Println(Mililiters(12.09248342))
}
