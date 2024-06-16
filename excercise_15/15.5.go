package main

import (
	"fmt"
	"sync"
	"time"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type StructMeteo struct {
	temp string
	mu   sync.RWMutex
}

func CreateStructMeteo(initialTemp string) *StructMeteo {
	return &StructMeteo{temp: initialTemp}
}

func (s *StructMeteo) ReadTemp() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.temp
}

func (s *StructMeteo) ChangeTemp(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.temp = v
}

func main() {
	meteo := CreateStructMeteo("10°C")
	var wg sync.WaitGroup

	//изменение температуры
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			time.Sleep(25 * time.Millisecond)
			temp := fmt.Sprintf("%+v°C", i)
			meteo.ChangeTemp(temp)
			fmt.Println("ChangeTemp:", temp)
		}
	}()

	//чтение температуры
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 50; i++ {
			time.Sleep(50 * time.Millisecond)
			temp := meteo.ReadTemp()
			fmt.Println("ReadTemp:", temp)
		}
	}()

	wg.Wait()
	fmt.Println("Finish.")
}
