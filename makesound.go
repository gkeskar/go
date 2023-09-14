package main

import "fmt"

type Whistle string

func (w Whistle) MakeSound() {
	fmt.Println("Tweet!")
}

type Horn string

func (h Horn) MakeSound() {
	fmt.Println("Honk!")
}

type NoiseMaker interface {
	MakeSound()
}

type Robot string

func (r Robot) MakeSound() {
	fmt.Println("Beep Beep")
}

func (r Robot) Walk() {
	fmt.Println("Powering legs")
}
func play(n NoiseMaker) {
	n.MakeSound()
}

func main() {
	var toy NoiseMaker
	toy = Whistle("Toyco Canary")
	play(toy)
	toy = Horn("Toyco Blaster")
	play(toy)
	play(Robot("Botco Ambler"))
	var noisemaker NoiseMaker = Robot("Botco Ambler")
	noisemaker.MakeSound()
	var robot Robot = noisemaker.(Robot)
	robot.Walk()
}
