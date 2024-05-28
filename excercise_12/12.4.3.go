/*
Вариант 3 - проверять указатель в методе Sing
*/

package main

import (
	"errors"
	"fmt"
)

type Bird3 interface{ Sing3() string }
type Duck3 struct{ voice string }

func (d *Duck3) Sing3() (string, error) {
	if d == nil {
		return "", errors.New("Ошибка пения!")
	}
	return d.voice, nil
}

func main() {
	var d *Duck3
	song, err := d.Sing3()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing3(b Bird3) (string, error) {
	if b != nil {
		return b.Sing3(), nil
	}
	return "", errors.New("Ошибка пения!")
}
