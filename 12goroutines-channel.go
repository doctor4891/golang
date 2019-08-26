/*
каналы
ch <- v    // Послать v в канал ch.
v := <-ch  // Получить из канала ch

создание канала ch := make(chan int)
создание буферизированного (вместимость) канала ch := make(chan intб)

По-сути результатами работы горутин можно обменивать с помощью каналов

ошибка deadlock возникает из за того что скрипт уже закончился а канал еще
готов на прием. Нужно или закрывать его или же установить его емкость при создании

каналы набивать нужно в горутине
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go countAnimal("dog", c)
	for {
		message, open := <-c
		if !open {
			break
		}
		fmt.Println(message)
	}
}

func countAnimal(animal string, c chan int) {
	for i := 1; i < 5; i++ {
		time.Sleep(time.Second)
		c <- i
	}
	close(c)
}
