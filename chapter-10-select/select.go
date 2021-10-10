package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "done1"
	}()

	go func() {
		time.Sleep(time.Second * 1)
		c2 <- "done2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("msg1=", msg1)
		case msg2 := <-c2:
			fmt.Println("msg2=", msg2)
		}
	}

}
