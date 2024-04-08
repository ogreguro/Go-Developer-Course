package main

import "fmt"

func main() {
	x := 16
	y := 3
	res := x / y
	ost := x % y

	fmt.Printf("%d / %d Результат: %d, остаток от деления: %d, тип результата: %T", x, y, res, ost, res)
}
