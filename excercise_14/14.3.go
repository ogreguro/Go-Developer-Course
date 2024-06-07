package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 4)

	go func() {
		ch <- "Привет"
		ch <- "буферизованный канал!"
	}()

	for i := 0; i < cap(ch); {
		select {
		case res := <-ch:
			fmt.Println(res)
		default:
			i++
			time.Sleep(time.Second)
		}
	}
}
