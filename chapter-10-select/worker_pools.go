package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//worker_pool()
	wait_groups()

}

func worker(id int, jobs <-chan int, result chan<- int) {
	for job := range jobs {
		fmt.Printf("worker %d start job %d \n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finish job %d  \n", id, job)
		result <- job
	}
}

func worker_pool() {
	const numJobs = 6
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs)
	for i := 1; i <= numJobs; i++ {
		<-results
	}
}

func worker_for_group(id int, group *sync.WaitGroup) {
	defer group.Done()

	fmt.Printf("time:%s, Worker %d starting\n", time.Now(), id)
	time.Sleep(time.Second * 8)
	fmt.Printf("time:%s, Worker %d done\n", time.Now(), id)
}

func wait_groups() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker_for_group(i, &wg)
	}
	wg.Wait()
	fmt.Println("all worker finished!!!!")

}
