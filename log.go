//простой пакет для вывода ошибок
//кроме простого вывода может вызвать fatal или panic
package main

import "log"

func main() {
	//простой вывод
	logPrint()

	logFatal()

	logPanic()
}

func logPrint() {
	log.Print("log print")
}

func logFatal() {
	log.Fatal("fatal error")
}

func logPanic() {
	log.Panic("panic error")
}
