package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//простой цикл который выводит числа в бинарном виде

	var i int64
	for i = 1; i <= 10; i++ {
		fmt.Println(i, strconv.FormatInt(i, 2))
	}
	/* result:
	1 1
	2 10
	3 11
	4 100
	5 101
	6 110
	7 111
	8 1000
	9 1001
	10 1010*/

	//вариант похожий на while в php
	i = 0
	for i < 10 {
		i++
		fmt.Println(i)
	}
	i = 0

	/*бесконечный цикл. Выйти можно только изнутри
	можно использовать например для написания сервера который будет постоянно слушать
	*/
	for {
		second := time.Second //1 секунда(константа)
		time.Sleep(time.Duration(second))
		fmt.Println(i)
		i++
		if i > 10 {
			break
		}
	}
}
