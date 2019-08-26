package main

import "fmt"

func main() {
	var a, b = "hello", "world"

	//стандартный вывод переменных
	fmt.Println(a, b)

	//print formated - форматированная печать
	//Форматирует строку - переносы, длина, добавление 0 и тд.
	//можно использовать из командной строки
	//использование: printf [-v var] format [arguments]
	//список аргументов смотреть в сети. https://www.php.net/manual/ru/function.sprintf.php
	fmt.Printf("%b\n", a+"\n"+b) //%b\n - означает интепретировать переносы,%d - число

	//Запись вывода в переменную save printf
	myStirng := fmt.Sprintf("%b\n", a+"\n"+b)
	fmt.Println(myStirng)

	//пишет заголовок в интернет, или же в файл (если
	// передать файловый дескриптор-типа объект описывающий файл)
	//
	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
}
