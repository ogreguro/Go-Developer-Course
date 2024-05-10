package main

import "fmt"

func main() {
	mp := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	mp["бегемот"] = 2
	fmt.Println(mp)
}
