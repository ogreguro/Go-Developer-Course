package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("Привет из дочерней горутины!")
		ch <- 0
	}()
	<-ch
}
