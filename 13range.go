//range - аналог foreach в php
//можно использовать и для каналов
package main

import "fmt"

func main() {
	mySlice := []int{1, 4, 6, 7, 4}

	var sum int

	for key, value := range mySlice {
		sum = sum + value
		fmt.Println(key)
	}
	fmt.Println(sum)

	/*ключи или значения можно не использовать ставя _ */
	for _, v := range mySlice {
		fmt.Println(v)
	}

	/*range для канала*/
	channel := make(chan string)
	/*сначала наполняем канал в горутине
	без горутины он не наполняется
	*/
	go func() {
		channel <- "dog"
		channel <- "cat"
		channel <- "ship"
		close(channel)
	}()

	/*для канала использовать только значение v. Ключа у канала нет*/
	for el := range channel {
		fmt.Println(el)
	}
}
