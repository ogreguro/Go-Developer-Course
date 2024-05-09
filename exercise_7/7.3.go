package main

import "fmt"

func main() {

	a := [...]string{"яблоко", "груша", "помидор", "абрикос"}
	a[2] = "персик"
	fmt.Println(a)
}
