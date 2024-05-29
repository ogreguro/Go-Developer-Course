package main

import (
	"encoding/json"
	"fmt"
)

type contract1 struct {
	Number   int
	SignDate *string
	Landlord string
	Tenat    string
}

func main() {
	str := `{"number":1,"landlord":"Остап Бендер","tenat":"Шура Балаганов"}`
	c := contract1{}
	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println("json.Unmarshal error:", err)
	}
	fmt.Printf("%+v", c)
	fmt.Println()
}
