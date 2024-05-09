package main

import "fmt"

func main() {
	var v int = 1
	var p *int = &v
	fmt.Println(v, p)
}
