package main

import "fmt"

func main() {
	ch := make(chan int)
	stop := make(chan struct{})
	go func() {
		<-ch
		stop <- struct{}{}
	}()

	close(ch) //закрываем канал для избежания блокировки
	<-stop
	fmt.Println("happy end")
}
