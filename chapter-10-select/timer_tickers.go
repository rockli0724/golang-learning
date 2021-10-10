package main

import (
	"bufio"
	"fmt"
	"net/http"
	"time"
)

func main() {

	//http_()
	//ticker()
	timer()
}

func http_() {
	resp, err := http.Get("http://gobyexample.com")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response.....", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
func timer() {

	timer := time.NewTimer(5 * time.Second)
	fmt.Println("time now is ", time.Now().Local())
	<-timer.C
	fmt.Println("time now is ", time.Now().Local())
	fmt.Println("timer1 has fired")

	newTimer := time.NewTimer(time.Second * 2)
	go func() {
		<-newTimer.C
		fmt.Println("time2 has fired")
	}()
	time.Sleep(3 * time.Second)
	stop := newTimer.Stop()
	if stop {
		fmt.Println("time2 has stop")
	}
}

func ticker() {
	newTicker := time.NewTicker(time.Second * 1)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-newTicker.C:
				fmt.Println("tick at ", t.Local())
			}
		}
	}()
	time.Sleep(3 * time.Second)
	newTicker.Stop()
	done <- true
	fmt.Println("ticker stop!!!")
}
