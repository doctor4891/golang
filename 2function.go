//Функции
package main

import "fmt"

//у функцию передаем переменные с обозначением типов
//и обозначаем какой тип хотим вернуть
func Calc(x, y int) int {
	return x + y
}

//функция может возвращать несколько результатов
//для каждого из них нужно указать тип
func Swap(a, b string) (string, string) {
	return a, b
}

//голый возврат
//функция может возвращать уже "поименованый" результат
func Naked(a, b string) (c string) {
	c = a + " " + b
	return
}

/*передача функции в функцию как значение*/
func MyFuncInFunc(myFunc func() int) int {
	return myFunc()
}


func main() {
	fmt.Println(Calc(1, 2))
	fmt.Println(Swap("Hello", "World"))
	fmt.Println(Naked("Naked", "digger"))

	/*функция как значение*/
	myFunc := func() int {
		return 1
	}
	fmt.Println(myFunc()) //1

	/*функцию можно передавать в функцию*/
	fmt.Println(MyFuncInFunc(myFunc)) //1

}

/*Функция с большой буквы может быть доступна в другом пакете пакет.Функция. Функция с
маленькой буквы может быть доступна только в этом пакете - такая приватность.*/

//////////основные моменты//////////////
// несколько результатов, голый возврат, функция как значение
