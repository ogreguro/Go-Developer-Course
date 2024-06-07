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
OUT:
	for {
		select {
		case res := <-ch:
			fmt.Println(res)
			if len(ch) == 0 {
				break OUT
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
