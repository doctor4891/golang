/*вся работа с сетью через этот пакет net*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	netDial()
}

/*подсоединяемся к серверу, отсылаем заголовок и получаем ответ*/
func netDial() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	log.Print(status) //2019/08/24 10:49:46 HTTP/1.0 200 OK

	//вывод полностью всего ответа
	// scanner := bufio.NewScanner(conn)
	//for scanner.Scan(){
	//log.Print(scanner.Text())
	//}
}
