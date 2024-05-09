package main

import "fmt"

type contacts struct {
	Addss, Phone string
}
type user struct {
	ID   int
	Name string
	contacts
}
type employee struct {
	ID   int
	Name string
	contacts
}

func main() {
	u := user{ID: 1, Name: "Ivan", contacts: contacts{Addss: "Ivan's home address", Phone: "111-111-111"}}
	e := employee{ID: 1, Name: "Peter", contacts: contacts{Addss: "Peter's home address", Phone: "222-222-222"}}
	fmt.Println(u.Addss, u.Phone, e.Addss, e.Phone)
}
