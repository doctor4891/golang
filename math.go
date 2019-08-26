package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//степень. Например 2 в 4 степени будет 16
	fmt.Println(math.Pow(2, 4)) //16

	//корень
	fmt.Println(math.Sqrt(9)) //3

	//случаное число в диапазоне до n
	fmt.Println(rand.Intn(10)) //1
}
