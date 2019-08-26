//понимаю как буфер обмена
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Scanner - штука которая может построчно пройтись по всем строкам или байтам файла
	// принимает поток ввода-вывода в буфер обмена
	//1. сначала нужен поток
	flow, _ := os.Open("testsFiles/source.txt")
	//сканером мы сканируем отдельно каждую страницу, так и здесь
	//scanner может построчно пройтись по строкам файла scanner.Scan()
	// и вывести каждую строку scanner.Text() или в байтах scanner.Bytes
	scanner := bufio.NewScanner(flow)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	//чтение потока из соединения
	//ReadString читает до ограничителя указанного в скобках
	// bufio.NewReader(conn).ReadString('\n')
}
