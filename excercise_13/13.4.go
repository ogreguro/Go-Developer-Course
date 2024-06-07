package main

import (
	"encoding/xml"
	"fmt"
)

type contract4 struct {
	Number   int
	SignDate string
	LandLord string
	tenat    string
}

type contracts4 struct {
	List []contract4 `xml:"contract"`
}

func main() {
	c := contract4{Number: 1, LandLord: "Остап Бендер", tenat: "Шура Балаганов"}
	cl := contracts4{}
	cl.List = append(cl.List, c)

	res, err := xml.Marshal(cl)
	if err != nil {
		fmt.Println("xml.Marshal error:", err)
	}

	fmt.Printf("%+v", c)
	fmt.Println()
	fmt.Printf("%+v", cl)
	fmt.Println()
	fmt.Printf("%v", string(res))
	fmt.Println()
}
