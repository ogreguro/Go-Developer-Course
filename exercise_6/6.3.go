package main

import "fmt"

type contract3 struct {
	ID           int
	number, Date string
}

func main() {
	c := contract3{ID: 1, number: `#000A\n101`, Date: "2024-01-31"}
	fmt.Println(c.show())
}

func (c *contract3) show() string {
	return fmt.Sprintf("Договор \u2116 %s от %s", c.number, c.Date)
}
