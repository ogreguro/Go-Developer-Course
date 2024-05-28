/*
Вариант 1 - проинициализировать структуру в main и передать ненулевой указатель
*/

package main

import (
	"errors"
	"fmt"
)

type Bird1 interface{ Sing1() string }
type Duck1 struct{ voice string }

func (d *Duck1) Sing1() string { return d.voice }

func main() {
	var d *Duck1 = &Duck1{"Кря!"}
	song, err := Sing1(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}
func Sing1(b Bird1) (string, error) {
	if b != nil {
		return b.Sing1(), nil
	}
	return "", errors.New("Ошибка пения!")
}
