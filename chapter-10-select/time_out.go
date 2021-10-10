package main

import (
	"fmt"
	"time"
)

func main() {
	//timeout()
	//nonBlockingChannelOperations()
	//closing_channels()
	rangeOverChannel()
}

func timeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	//select 可以起到阻塞等待的效果
	// https://studygolang.com/articles/7203
	select {
	case r1 := <-c1:
		fmt.Println(r1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	select {
	case r2 := <-c2:
		fmt.Println(r2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2 ")
	}

}

func nonBlockingChannelOperations() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages:
		fmt.Println("received msg", msg)
	default:
		fmt.Println("no message received..")
	}
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("send msg", messages)
	default:
		fmt.Println("no msg sent")
	}
	select {
	case msg := <-messages:
		fmt.Println("receive msg", msg)
	case sig := <-signals:
		fmt.Println("receive sig", sig)
	default:
		fmt.Println("no act")
	}

}

func closing_channels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("receive job ", j)
			} else {
				fmt.Println("receive  all jobs ", j)
				done <- true
				return
			}
		}
	}()

	for i := 1; i <= 3; i++ {
		time.Sleep(3 * time.Second)
		jobs <- i
		fmt.Println("send job ", i)
	}
	close(jobs)
	fmt.Println("send all jobs")
	<-done
}

func rangeOverChannel() {
	s := make(chan string, 2)
	s <- "hell1"
	s <- "hello2"
	close(s)
	for s2 := range s {
		fmt.Println(s2)
	}

}
