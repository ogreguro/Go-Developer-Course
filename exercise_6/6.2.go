package main

import "fmt"

type contract2 struct {
	ID           int
	number, Date string
}

func main() {
	c := contract2{ID: 1, number: "#000A\t101", Date: "2024-01-31"}
	fmt.Println(c.show())
}

func (c *contract2) show() string {
	return fmt.Sprintf("ID: %d, Number: %s, Date:%s", c.ID, c.number, c.Date)
}
