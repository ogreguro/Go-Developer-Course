package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)

	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v) // Болтливая горутина
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()

	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()

	//пока в канал stop не поступило значение, горутина 1 будет вести счет каждый раз, пока горутина 2 будет записывать значения в ch.
	//в горутине 2 идет проверка ch == nil, что заставит горутину 2 начать следующий цикл проверок при удовлетворении этого условия.
	//после 5 секунд ожидания в канал stop поступят знеачения и обе горутины остановят работу
	ch = nil

	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)

	fmt.Println("завершение работы главной горутины")
}
