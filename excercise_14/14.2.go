package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Привет из дочерней горутины!"
	}()
	fmt.Println(<-ch)
}
