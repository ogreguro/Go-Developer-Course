package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := []int{}
	c = append(c, a...)
	c = append(c, b...)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
