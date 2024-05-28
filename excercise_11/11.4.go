package main

import (
	"errors"
	"fmt"
)

/*
Нужно создать, используя оборачивание, ошибку «ошибка3:ошибка2:ошибка1».
Также нужно создать свою ошибку в виде структуры myFirstError, которая обязательно должна иметь метод Error() string.
Необходимо убедиться, что в созданной цепочке ошибок не было ошибки типа myFirstError
*/
type myFirstError struct {
	status  int
	message string
}

func (e myFirstError) Error() string { return e.message }

func main() {
	wanted_error := myFirstError{
		status:  1,
		message: "искомая ошибка",
	}
	fmt.Println(wanted_error)
	err := errors.New("ошибка1")
	err = fmt.Errorf("ошибка2:%w", err)
	err = fmt.Errorf("ошибка3:%w", err)
	fmt.Println(err)
	if !errors.As(err, new(myFirstError)) {
		fmt.Println("ошибки \"myFirstError\" не было")
	}
}
