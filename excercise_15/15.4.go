package main

import (
	"fmt"
	"sync"
)

func start() {
	fmt.Println("func start()")
}

func main() {
	var so sync.Once
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			i := i
			so.Do(start)
			fmt.Println("goroutine child:", i+1)
		}()

	}
	wg.Wait()
	fmt.Println("goroutine main")
}
