package main

import "fmt"

type Bird interface {
	Fly()
}
type Duck struct{}

func (d Duck) Fly() {
	fmt.Println("Утка - Умею летать!")
}

func (d Duck) Swim() {
	fmt.Println("Утка - Умею плавать!")
}

type Sparrow struct{}

func (s Sparrow) Fly() {
	fmt.Println("Воробей - Умею летать!")
}

func main() {
	var d, s Bird
	d = Duck{}
	Do(d)
	s = Sparrow{}
	Do(s)
}
func Do(b Bird) {
	b.Fly()
	//здесь нужно дописать код
	c, ok := b.(Duck)
	if ok {
		c.Swim()
	}
}
