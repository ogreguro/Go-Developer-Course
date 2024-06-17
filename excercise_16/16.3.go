package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(2*time.Second))
	defer cancel()

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("стоп горутина:", i+1, ctx.Err())
					return
				default:
					fmt.Println("работаю", i+1)
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()
	}
	
	wg.Wait()
}
