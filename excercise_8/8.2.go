package main

import "fmt"

func main() {
	mp := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	k := "слон"
	m, ok := mp[k]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)", k, m, ok)
	k = "бегемот"
	m, ok = mp[k]
	fmt.Printf("\nЖивотное: %s, количество: %d (есть в карте: %v)", k, m, ok)
	k = "осьминог"
	m, ok = mp[k]
	fmt.Printf("\nЖивотное: %s, количество: %d (есть в карте: %v)", k, m, ok)
}
