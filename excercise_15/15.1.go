package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			i := i
			defer wg.Done()
			fmt.Println("горутина: ", i+1)
		}()
	}
	wg.Wait()
}
