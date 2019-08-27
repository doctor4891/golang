///////////////		Интерфейсное значение 		/////////////////////
// Сначала переменной присваиваем интерфейс, затем тип. И потом этот тип жестко реализует все что написано в интерфейсе

package main

import "fmt"

type Car interface {
	BuyCar() float32
}

type PriceHoliday struct {
	price float32
	discount float32
}

type PriceUsual struct {
	price float32
}

/*расчет цены на авто со скидкой*/
func (priceData PriceHoliday) BuyCar()float32  {
	return priceData.price*(1-priceData.discount/100)
}
/*расчет цены на авто без скидки*/
func (priceData PriceUsual) BuyCar() float32  {
	return priceData.price
}

func main()  {
	/*интерфейсное значение - полюбасу должно крякнуть-реализовать функцию интерфейса на покупку машины*/
	var honda Car
	honda = PriceHoliday{10000,10}//в honda записывается и интерфейс с его методом и тип со значениями
	fmt.Println(honda.BuyCar())//9000

	honda = PriceUsual{10000}//здесь уже реализация по другому типу, но тот же интeрфейс
	fmt.Println(honda.BuyCar())//10000
}