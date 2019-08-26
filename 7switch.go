package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//псевдослучайное число в диапазоне n
	a := rand.Intn(2) //1

	/*switch с условием*/
	switch a {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	//	switch без условия (switch true)
	switch {
	case a == 1:
		fmt.Println("one")
	case a == 2:
		fmt.Println("two")
	}
}
