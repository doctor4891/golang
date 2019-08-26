///////////////////// КАРТА MAP ///////////////////////
package main

import "fmt"

func main() {
	/*создание карты*/
	myMap := make(map[string]string)
	fmt.Println(myMap) //map[]

	/*добавление значения в карту*/
	myMap["animal"] = "dog"

	fmt.Println(myMap)           //map[animal:dog]
	fmt.Println(myMap["animal"]) //dog

	/*обход всех значений карты*/
	for k, v := range myMap {
		fmt.Println(k, v) //animal dog
	}

	/*проверка наличия ключа*/
	elem, exist := myMap["animal1"]
	if exist == false {
		fmt.Println("element not exist")
	} else {
		fmt.Println(elem)
	}

	/*удалить элемент*/
	key := "animal" //ключ по которому удаляем
	delete(myMap, key)
	fmt.Println(myMap) //map[]
}
