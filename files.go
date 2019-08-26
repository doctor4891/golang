//работа с файлами

//полезный источник https://golangbot.com/write-files/
//термин "файловый дескриптор" это собственно объект описывающий файл
//пример myFile, _ := os.OpenFile("testsFiles/file1.txt", os.O_APPEND|os.O_WRONLY, 0666)
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//скопировать содержание одного файла в другой файл
	//copyFile("testsFiles/source.txt","testsFiles/destination.txt")

	//считать файл построчно
	//readFileLineByLine()

	//запись в файл
	//writeToFile()

	/*запись вконец файла*/
	appendToTheEndOfFile()
}

/*копирование одного файла в другой*/
func copyFile(src, dst string) (written int64, err error) {
	//возвращает EOF - end of file, то есть количество символов по сути

	//открытие файла
	source, err := os.Open(src)
	if err != nil {
		fmt.Println("source open problem")
		return
	}
	defer source.Close()
	//создание или открытие другого файла
	destination, err := os.Create(dst)
	if err != nil {
		fmt.Println("destination create problem")
		return
	}
	defer destination.Close()
	//копирование
	written, err = io.Copy(destination, source)
	if err != nil {
		fmt.Println("copy problem")
	}
	fmt.Println(written)

	return
}

/*чтение файла построчно*/
func readFileLineByLine() {
	myFile, err := os.Open("testsFiles/socks.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()

	scanner := bufio.NewScanner(myFile)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*Запись в файл*/
func writeToFile() {
	myFile, err := os.Create("testsFiles/file1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer myFile.Close()
	writer, err := myFile.WriteString("mystring")
	if err != nil {
		log.Fatal("filed write ", err)
	}
	log.Println(writer)
}

func appendToTheEndOfFile() {
	myFile, _ := os.OpenFile("testsFiles/file1.txt", os.O_APPEND|os.O_WRONLY, 0666)

	defer myFile.Close()

	fmt.Fprintf(myFile, "Sting number two")

}
