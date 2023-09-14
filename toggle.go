package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type Toggable interface {
	toggle()
}

func main() {
	s := Switch("off")
	var t Toggable = &s
	t.toggle()
	t.toggle()
}
