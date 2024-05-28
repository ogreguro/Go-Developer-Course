package main

import "fmt"

func main() {
	a := 1
	do(a)
	fmt.Println("выполнение продолжается")

	b := "1"
	do(b)
	fmt.Println("выполнение успешно завершено")
}
func do(v any) {
	a := 10
	b, ok := v.(int)
	if !ok {
		panic("не удалось привести значение v к int")
	}
	a += b
	fmt.Println(a)
}
