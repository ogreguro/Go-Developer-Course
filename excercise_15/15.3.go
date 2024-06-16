package main

import (
	"fmt"
	"sync"
)

func main() {
	var n int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			n++
		}()
	}
	wg.Wait()
	fmt.Println(n)
}
