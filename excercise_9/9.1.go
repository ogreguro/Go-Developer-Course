package main

import "fmt"

func main() {
	// Тестирование функции
	fmt.Println("апельсинов:", fruitMarket("апельсины"))
	fmt.Println("груш:", fruitMarket("груши"))
	fmt.Println("ананасов:", fruitMarket("ананасы"))
}

func fruitMarket(fruitName string) int {
	fruits := map[string]int{
		"апельсины": 5,
		"яблоки":    3,
		"сливы":     1,
		"груши":     0,
	}

	quantity, ok := fruits[fruitName]
	if !ok {
		fmt.Printf("Фрукт '%s' на рынке отсутствует\n", fruitName)
	}
	return quantity
}
