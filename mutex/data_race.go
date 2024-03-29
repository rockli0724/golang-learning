package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var n int

func main() {
	var wg sync.WaitGroup
	nCPU := runtime.NumCPU()
	nIter := 10
	for i := 0; i < nCPU; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < nIter; j++ {
				n++
				time.Sleep(time.Microsecond * 10)

			}
		}()
	}
	wg.Wait()
	fmt.Printf("n is: %d, expected: %d\n", n, nCPU*nIter)
}
