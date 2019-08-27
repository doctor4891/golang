/////////////// 	ИНТЕРФЕЙСЫ    ///////////////
// по сути это наборы методов
// они нужны просто для описания, удобства организации, философии проекта
// в го реализуют "утиное типизировани" - если крякает и плавает то это утка
// поэтому сначала пишем структуры, затем методы c большой буквы, а только потом выходим на интерфейсы
// если они нужны конечно же
package main

import "fmt"

type Auth interface {
	Login()
	Logout()
}

/*тип реализует все методы интерфейса*/
type User struct {
}

func (User) Login(){
	fmt.Println("Login")
}

func (User) Logout()  {
	fmt.Println("Loguot")
}

func main()  {
	newUser := User{}
	newUser.Login()
	newUser.Logout()
}

