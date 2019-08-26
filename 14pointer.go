package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "bull"
	fmt.Println(a) //bull
	/*адрес ячейки памяти a*/
	b := &a
	fmt.Println(b) //0xc00000e1e0
	//todo тип ссылка на строку, хотя по сути это тип uintptr Почему?

	fmt.Println(reflect.TypeOf(b)) //*string
	//указатель на значение
	fmt.Println(*b) //bull
}

/*
1. Память поделена на ячейки.
2. Переменная это псевдоним адреса этой ячейки
	По сути a := bull и 0xc00000e1e0 := bull одно и тоже

3. &a это еще одна ячейка в памяти в которой хранится только адрес ячейки a
4. * извлечь значение по записаному адресу
*/
