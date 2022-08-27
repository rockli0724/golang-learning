package main

import (
	"fmt"
	"sync"
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
	//channel_directions()
	//ch := make(chan int, 64)
	//go Producer(3, ch)
	//go Producer(5, ch)
	//go Consumer(ch)
	////time.Sleep(1 * time.Second) //1.通过睡眠方式
	//
	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//fmt.Printf("quit (%v)\n", <-sig)
	//

	test()
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

func Producer(facotr int, out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * facotr
	}
}
func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println("consumer: ", v)
	}
}

type httpPkg struct {
}

func (httpPkg) Get(url string) {
	time.Sleep(3 * time.Second)
	fmt.Println("pong....", url)
}

var http httpPkg
var myStr string

type myS string

func (receiver myS) name() {
	return

}

//多goroutine 请求
func test() {
	now := time.Now()
	var wg sync.WaitGroup
	var urls = []string{
		"https://www.baidu.com",
		"https://www.jd.com",
		"https://www.sougou.com",
	}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			http.Get(url)
		}(url)
	}
	wg.Wait()
	fmt.Println("cost time =", time.Since(now).Seconds())

	str := "Go爱好者"
	fmt.Printf("The string: %q\n", str)
	fmt.Printf("  => runes(char): %q\n", []rune(str))
	fmt.Printf("  => runes(hex): %x\n", []rune(str))
	fmt.Printf("  => bytes(hex): [% x]\n", []byte(str))
}
