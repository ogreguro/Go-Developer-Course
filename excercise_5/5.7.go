package main

import "fmt"

type square2 int

func main() {
	var s square2 = 30
	fmt.Println(s + 15)
	//или
	s += 15
	fmt.Println(s)
}
