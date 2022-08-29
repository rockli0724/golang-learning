package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	semaphoreSize      = 4
	mu                 sync.Mutex
	totalTasks         int
	cuConcurrentTasks  int
	maxConcurrentTasks int
)

func timeConsumingTask() {
	mu.Lock()
	totalTasks++
	cuConcurrentTasks--
	if cuConcurrentTasks > maxConcurrentTasks {
		maxConcurrentTasks = cuConcurrentTasks
	}
	mu.Unlock()

	//real system operation
	time.Sleep(time.Second * 5)
	mu.Lock()
	cuConcurrentTasks--
	mu.Unlock()
}

func main() {
	//sem := make(chan struct{}, semaphoreSize)
	//var wg sync.WaitGroup
	//now := time.Now()
	//for i := 0; i < 32; i++ {
	//	sem <- struct{}{}
	//	wg.Add(1)
	//
	//	go func() {
	//		defer wg.Done()
	//		timeConsumingTask()
	//		<-sem
	//	}()
	//}
	//wg.Wait()
	//fmt.Printf("totalTask:%d \n", totalTasks)
	//fmt.Printf("maxConCurrentTasks :%d \n", maxConcurrentTasks)
	//fmt.Printf("totalCost :%v \n", time.Since(now))
	unbuffered()
	buffered()

}

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			time.Sleep(2 * time.Second)
		} else {
			time.Sleep(time.Second)
		}
		ch <- i
	}
}

func consumer(ch chan int) {
	total := 0
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			time.Sleep(2 * time.Second)
		} else {
			time.Sleep(time.Second)
		}
		total += <-ch
	}
}

func unbuffered() {
	now := time.Now()
	intsCh := make(chan int)
	go producer(intsCh)
	consumer(intsCh)
	fmt.Printf("unfuffered verion took %s \n", time.Since(now))
}

func buffered() {
	now := time.Now()
	intsCh := make(chan int, 5)
	go producer(intsCh)
	consumer(intsCh)
	fmt.Printf("fuffered verion took %s \n", time.Since(now))
}
