package main

import (
	"encoding/xml"
	"fmt"
)

type contract3 struct {
	Number   int    `xml:"number"`
	SignDate string `xml:"sign_date"`
	Landlord string `xml:"landlord"`
	Tenat    string `xml:"tenat"`
}
type contracts3 struct {
	List []contract3 `xml:"contract"`
}

func main() {
	str :=
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>" +
			"<contracts>" +
			"<contract>" +
			"<number>1</number>" +
			"<sign_date>2023-09-02</sign_date>" +
			"<landlord>Остап</landlord>" +
			"<tenat>Шура</tenat>" +
			"</contract>" +
			"<contract>" +
			"<number>2</number>" +
			"<sign_date>2023-09-03</sign_date>" +
			"<landlord>Бендер</landlord>" +
			"<tenat>Балаганов</tenat>" +
			"</contract>" +
			"</contracts>"
	fmt.Println(str)
	c := contracts3{}
	err := xml.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println("xml.Unmarshal error:", err)
	}
	fmt.Printf("%+v", c)
}
