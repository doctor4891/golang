package main

import (
	"fmt"
	"strings"
)

func main() {
	/*объединение среза в строку
	/*аналог implode в php*/
	joinString()
}

/*объединение среза в строку*/
/*аналог implode в php*/
func joinString() {
	myStrings := []string{"Hello", "World"}

	joinedString := strings.Join(myStrings, "***")

	fmt.Println(joinedString) //Hello***World
}
