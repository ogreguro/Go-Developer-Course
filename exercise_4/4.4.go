package main

import "fmt"

func main() {
	hello44()
}

func hello44() func() {
	f := func() {
		fmt.Println("Hello, Go!")
	}
	return f
}
