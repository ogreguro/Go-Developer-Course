package main

import "fmt"

func main() {
	var v int = 1
	var p *int = &v
	fmt.Println(v, p)
	change(p)
	fmt.Println(v, p)
}

func change(p *int) {
	*p = 2
}
