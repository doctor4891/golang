package main

import (
	"fmt"
	"time"
)

func main() {

	//текущее время полностью
	fmt.Println(time.Now()) // 2019-08-22 16:07:56.890166708 +0300 EEST m=+0.000218288

	//день недели
	fmt.Println(time.Now().Weekday()) //например Thursday

	/*пауза в 1 секунду*/
	time.Sleep(time.Second)

	//текущая минута
	fmt.Println(time.Now().Minute()) //например 32

}
