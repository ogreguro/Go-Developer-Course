package main

import "fmt"

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	go func() {
		<-ch
		stop <- struct{}{}
	}()
	close(ch)
	<-stop
	fmt.Println("happy end")
}
