////////////// МЕТОДЫ //////////////////////
// это функции для типов
package main

import "fmt"

/*метод добавляет значение по указателю
если указатель не указать, то просто будут данные внутри функции
и чтобы они вышли наружу нужно будет их return и в основной функции установить
значение. А это дополнительные расходы на копирование.
*/
func (allUsers *usersCount) addUserItem() {
	*allUsers++
}

func (user *userData) changeUserStatus() {
	user.userStatus = "registered"
}

/*количество пользователей сайта*/
type usersCount int

/*данные конкретного пользователя*/
type userData struct {
	userId     int
	userName   string
	userStatus string
}

func main() {
	//легенда: добавление пользователя на сайт, изменение его статуса
	allUsers := usersCount(10)
	allUsers.addUserItem()
	fmt.Println(allUsers)

	user := userData{1, "Bob", "guest"} //& можно не ставить

	user.changeUserStatus()

	fmt.Println(user.userStatus) //registered

}
