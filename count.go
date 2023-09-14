package main

import "fmt"

func count(start int, end int) {
	fmt.Printf("count(%d, %d) called\n", start, end)
	fmt.Println(start)
	if start < end {
		count(start+1, end)
	}
}

func main() {
	count(2, 5)
}
