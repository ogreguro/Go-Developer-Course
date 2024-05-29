package main

import (
	"encoding/json"
	"fmt"
)

type contract2 struct {
	Number   int    `json:"number"`
	SignDate string `json:"sign_date,omitempty"`
	Landlord string `json:"landlord"`
	tenat    string //неэкспортируемое поле
}

func main() {
	c := contract2{
		Number:   2,
		Landlord: "Остап",
		tenat:    "Шура",
	}

	fmt.Printf("%+v", c)
	fmt.Println()
	res, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json.Marshal error:", err)
	}
	fmt.Printf("%v", string(res))
	fmt.Println()
}
