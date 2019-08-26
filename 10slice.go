//Срезы
package main

import "fmt"

func main() {
	/*создание среза из массива*/
	/*массив*/
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArray) //[1 2 3 4 5]

	/*создаем срез (слайс)*/
	mySlice := myArray[0:4]
	fmt.Println(mySlice) //[1 2 3 4]

	/*или же другими способами*/
	mySlice1 := myArray[:]  //на ввесь массив
	fmt.Println(mySlice1)   //[1 2 3 4 5]
	mySlice2 := myArray[3:] //на другую часть массива
	fmt.Println(mySlice2)   //[4 5]

	/*еще один способ создание среза,
	  с помощью make*/
	mySlice3 := make([]int, 5)
	fmt.Println(mySlice3) //[0 0 0 0 0]

	//изменение значения в срезе
	mySlice3[2] = 3
	fmt.Println(mySlice3[2]) //3

	//добавление значения в конец среза

	mySlice3 = append(mySlice3, 6)
	fmt.Println(mySlice3) //[0 0 3 0 0 6]

	//длина среза - количество элементов в нем
	fmt.Println(len(mySlice3)) //6

	//вместимость среза - есть размер ячейки памяти выделенный для этого среза
	//если срез будет 11 элементов - выделят в памяти запас в 22 элемента
	fmt.Println(cap(mySlice3)) //10

	/*слайс как указатель на массив
	  при измененеии элемента слайса - изменяется этот же элемент массива
	*/
	sliceAsPointerToArray()

	/*срез среза*/
	sliceSlices := [][]int{mySlice1, mySlice2, mySlice3}
	fmt.Println(sliceSlices) //[[1 2 3 4 5] [4 5] [0 0 3 0 0 6]]
}

/*слайс как указатель на массив*/
func sliceAsPointerToArray() {
	/*создадим массив*/
	myArray := [5]int{1, 2, 3, 4, 5}

	/*создадим срез из этого массива*/
	mySlice := myArray[2:4]
	fmt.Println(mySlice) //[3 4]

	/*изменим элемент слайса*/
	mySlice[0] = 100
	/*этот же элемент изменился и в массиве*/
	fmt.Println(myArray) //[1 2 100 4 5]
}
