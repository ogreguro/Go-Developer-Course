package main

import "fmt"

func main() {
	defer fmt.Println("завершение работы")
	f := func() {
		fmt.Println("Hello, Go!")
	}
	f()
}
