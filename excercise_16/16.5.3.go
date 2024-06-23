package main

import (
	"fmt"
	"sync"
	"time"
)

//доп вариант с использованием потокобезопасного доступа к общей переменной через мутексы

func main() {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	rc := 10
	stop := false

	for i := 0; i < rc; i++ {
		wg.Add(1)
		go func(id int) {
			for {
				fmt.Printf("сложные вычисления горутины: %d\n", id+1)
				time.Sleep(1 * time.Second)

				mu.RLock()
				if stop {
					mu.RUnlock()
					fmt.Printf("stop горутина: %d\n", id+1)
					return
				}
				mu.RUnlock()
			}
		}(i)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")

		mu.Lock()
		stop = true
		mu.Unlock()
	}()

	wg.Wait()
}
