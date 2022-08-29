package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	cache map[int]int
	mu    sync.Mutex
)

func expensiveOperation(n int) int {
	time.Sleep(time.Second)
	return n * n
}

func getCache(n int) int {

	mu.Lock()
	v, ok := cache[n]
	mu.Unlock()
	if ok {
		return v
	}

	v = expensiveOperation(n)
	mu.Lock()
	cache[n] = v

	mu.Unlock()
	return v
}

func accessCache() {
	var total int
	for i := 0; i < 5; i++ {
		n := getCache(i)
		total += n
	}
	fmt.Printf("total %d \n", total)
}

//func main() {
//	cache = make(map[int]int)
//	go accessCache()
//	accessCache()
//}
