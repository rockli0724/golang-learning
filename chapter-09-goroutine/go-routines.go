package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, " : ", i)
	}
}

func main() {
	//basic_goroutine()
	//channel_buffer()
	//channel_synchronization()
	channel_directions()
}
func basic_goroutine() {
	f("direct")

	go func(msg string) {
		fmt.Printf("go func (%v)", msg)
	}("going")
	go f("goroutine")

	time.Sleep(time.Second)
	fmt.Println("done()")
}

func channel_buffer() {
	message := make(chan string, 4)
	message <- "ping"
	message <- "ping2"
	message <- "ping3"
	message <- "ping4"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}

// 通道同步
func channel_synchronization() {
	done := make(chan bool, 1)
	go work(done)
	<-done //同步阻塞获取channel中的值，保证同步结束的关键之处
	fmt.Println("finish.....")
}

func work(done chan bool) {
	fmt.Print("working.....")
	time.Sleep(time.Second * 3)
	fmt.Println("done....")
	done <- true //将完成的标志 塞进channel,保证主线程能够正常获取成功
}

func channel_directions() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "messages....")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings //将ping的消息取出来给msg
	pongs <- msg   //msg 消息再转给pongs
}
