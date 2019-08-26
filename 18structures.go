/////////////////////	STRUCT СТРУКТУРЫ	//////////////////////
//структуры это коллекция полей
package main

import "fmt"

/*структура это что-то типа каркаса для будущего "объекта"*/
type myStruct struct {
	myNumber   int
	myWord     string
	myDecision bool
}

func main() {
	fmt.Println(myStruct{}) //{0  false} - поля по умолчанию

	/*создаем объект на основе структуры - "натягиваем" данные на каркас*/
	coolObject := myStruct{1, "World", true}
	fmt.Println(coolObject) //{1 World true}

	/*можно вывести только то что интересует*/
	fmt.Println(coolObject.myWord) //World

	/*можно использователь указатели*/
	superObject := &coolObject
	fmt.Println(superObject.myWord) //World * Не не обязательно ставить - упростили
	superObject.myWord = "Earth"
	fmt.Println(coolObject.myWord) //Earth

}
