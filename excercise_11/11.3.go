package main

import (
	"errors"
	"fmt"
)

func main() {
	err_1 := errors.New("ошибка1")
	err := fmt.Errorf("ошибка2:%w", err_1)
	err = fmt.Errorf("ошибка3:%w", err)
	fmt.Println(err)
	fmt.Println("ошибка \"ошибка1\" была:", errors.Is(err, err_1))
}
