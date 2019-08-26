package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "sleep 1 second"
			time.Sleep(time.Second * 1)
		}

	}()

	go func() {
		for {
			ch2 <- "sleep 2 seconds"
			time.Sleep(time.Second * 2)
		}

	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}
