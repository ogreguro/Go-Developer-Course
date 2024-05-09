package main

import "fmt"

func main() {

	var a []int
	a = make([]int, 10, 10)
	a = append(a, 4, 1, 8, 9)
	fmt.Println(a)
}
