package main

import "fmt"

type contract struct {
	ID           int
	number, Date string
}

func main() {
	c := contract{ID: 1, number: `#000A\n101`, Date: "2024-01-31"}
	fmt.Println(c.show())
}

func (c *contract) show() string {
	return fmt.Sprintf("ID: %d, Number: %s, Date:%s", c.ID, c.number, c.Date)
}
