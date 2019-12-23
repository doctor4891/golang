/*mutex это блокировщик для горутин общего доступа к ресурсу*/
/*https://metanit.com/go/tutorial/7.6.php*/
package main

import (
	"fmt"
	"sync"
)

func counter(channel chan bool, gonum int, mutex *sync.Mutex)  {
	/*блокировка пока одна горутина работает с итератором*/
	mutex.Lock()
	for i:=1; i<10; i++{
		fmt.Println("goruoutine", gonum, "---", i)
	}
	mutex.Unlock()
	channel<-true
}

func main() {
	channel := make(chan bool)
	var mymutex sync.Mutex

	for i:=1; i<10; i++ {
		go counter(channel, i, &mymutex)
	}

	for i:=1; i<10; i++ {
		<-channel
	}
}