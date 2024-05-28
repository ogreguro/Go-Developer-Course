/*
Вариант 2 - проверять указатель на предмет нулевого значения в функции Sing
*/

package main

import (
	"errors"
	"fmt"
)

type Bird2 interface{ Sing2() string }
type Duck2 struct{ voice string }

func (d *Duck2) Sing2() string { return d.voice }

func main() {
	var d *Duck2
	song, err := Sing2(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}
func Sing2(b Bird2) (string, error) {
	duck, ok := b.(*Duck2)
	if b != nil && duck != nil && ok {
		return b.Sing2(), nil
	}
	return "", errors.New("Ошибка пения!")
}
