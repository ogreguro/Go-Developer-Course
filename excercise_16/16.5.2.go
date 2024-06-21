package main

import (
	"fmt"
	"sync"
	"time"
)

//доп. вариант
//закрытие канала избавляет от лишних горутин, пишущих значения в каналы в управляющей горутине
//вместо проверки на наличие значения в канале в рабочих горутинах осуществляется проверка на открытость - попытка чтения из него

func main() {
	var wg sync.WaitGroup
	const rc int = 10

	var channels [rc]chan struct{}

	for i := 0; i < rc; i++ {
		wg.Add(1)
		channels[i] = make(chan struct{})
		go func(i int, ch chan struct{}) {
			defer wg.Done()
		OUT:
			for {
				select {
				case _, ok := <-ch:
					if !ok {
						fmt.Printf("stop горутина: %d\n", i+1)
						break OUT
					}
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
			close(ch)
		}
	}()
	wg.Wait()
}
