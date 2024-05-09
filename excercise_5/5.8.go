package main

import "fmt"

type square3 int

func main() {
	var s square3 = 34
	s += 10
	fmt.Printf("%d m\u00B2", s)
}
