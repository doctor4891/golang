/*
многопоточность - go рутины, или go программы
  множество маленьких программок могут выполнятся одновременно
  выполняются в одном пространстве памяти
  запуск go func()

sync.Wait - для того чтобы можно было ждать пока горутины завершаться
и получить результат

*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*ожидатель, когда горутины завершаться*/
	var wg sync.WaitGroup
	/*добавляем 1 горутину в ожидание*/
	wg.Add(1)
	go func() {
		num("cow")
		/*удаляем из ожидателя*/
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		num("rat")
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		num("cat")
		wg.Done()
	}()
	wg.Wait()
}

func num(animal string) {
	for i := 1; i < 5; i++ {
		fmt.Println(i, animal)
		time.Sleep(time.Second)
	}
}
