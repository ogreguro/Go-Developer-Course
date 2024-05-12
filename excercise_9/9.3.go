package main

import "fmt"

func main() {
	checkFood("апельсин")
	checkFood("помидор")
	checkFood("клубника")
}

func checkFood(foodName string) {
	switch foodName {
	case "груша", "яблоко", "апельсин":
		fmt.Println("это фрукт")
	case "тыква", "огурец", "помидор":
		fmt.Println("это овоц")
	default:
		fmt.Println("что-то странное")
	}
}
