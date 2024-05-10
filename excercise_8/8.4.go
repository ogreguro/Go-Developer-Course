package main

import "fmt"

type animal2 struct{}

func main() {
	mp := map[string]animal2{
		"слон":    animal2{},
		"бегемот": animal2{},
		"носорог": animal2{},
		"лев":     animal2{},
	}
	mp["выдра"] = animal2{}
	fmt.Println(mp)
}
