package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var cntr int32
	routines := 10
	cycles := 100
	for i := 0; i < routines; i++ {
		go func() {
			for i := 0; i < cycles; i++ {
				atomic.AddInt32(&cntr, 1)
			}
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("значение счетчика =", cntr)
}
