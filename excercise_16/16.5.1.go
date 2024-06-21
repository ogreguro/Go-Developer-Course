package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	const rc int = 10

	var channels [rc]chan int

	for i := 0; i < rc; i++ {
		wg.Add(1)
		channels[i] = make(chan int)
		go func(i int, ch chan int) {
			defer wg.Done()
		OUT:
			for {
				select {
				case <-ch:
					fmt.Printf("stop горутина: %d\n", i+1)
					break OUT
				default:
					fmt.Printf("сложные вычисления горутины: %d\n", i+1)
					time.Sleep(1 * time.Second)
				}
			}
		}(i, channels[i])
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		for _, ch := range channels {
			go func(ch chan int) {
				ch <- 1
			}(ch)
		}
	}()
	wg.Wait()
}
