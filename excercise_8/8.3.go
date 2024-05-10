package main

import "fmt"

type animal1 struct{}

func main() {
	mp := map[string]animal1{
		"слон":    animal1{},
		"бегемот": animal1{},
		"носорог": animal1{},
		"лев":     animal1{},
	}
	delete(mp, "бегемот")
	fmt.Println(mp)
}
