package main

import "fmt"

//panic - 1. останавливает скрипт текущей функции
// 		  2. defer сработает
//		  3. функция возвращает panic и если нет recover паникует эта тоже функция и тд.

//recover - 1. Если нет паник то recover() имеет значение nil
//			2. Если есть паника то recover() имеет значение паники
//			3. Таким образом с помощью defer и recover() можно отловить панику и продолжить выполнение
func main() {
	defer func() {
		/*отлов паники*/
		if panicRes := recover(); panicRes != nil {
			fmt.Println("recovered", panicRes)
		}
	}()

	simplePanic()
}

func simplePanic() {
	/*defer выполнится в любом случае*/
	defer fmt.Println("Panicking")
	/*паникуем*/
	panic(fmt.Sprintf("%s", "its a panic"))
}
