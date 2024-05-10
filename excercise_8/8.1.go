package main

import "fmt"

type animal struct{}

func main() {
	m := map[string]animal{
		"слон":    animal{},
		"бегемот": animal{},
		"носорог": animal{},
		"лев":     animal{},
	}
	fmt.Println(m)
}
