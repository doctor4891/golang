//os - operating system
// пакет для взаимодействия с операционной системой
//что-то по типу командной строки в linux
package main

import "os"

func main() {
	//создание файла
	os.Create("testsFiles/file1.txt")
	//открытие файла на чтение
	os.Open("testsFiles/file1.txt")
}
