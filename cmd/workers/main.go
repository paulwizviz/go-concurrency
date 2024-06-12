package main

import (
	"fmt"
)

type Signal struct {
	Status int
}

func worker(id int, s <-chan Signal, r chan<- string) {
	for sig := range s {
		r <- fmt.Sprintf("Worker: %d, Signal Status: %d", id, sig.Status)
	}
}

func main() {

	// We match the number of jobs and results
	numJobs := 10
	jobs := make(chan Signal, numJobs)
	results := make(chan string, numJobs)

	// We restrict the number of workers and Goroutine to
	// 20. We use this routine to restrict
	for w := 0; w < 20; w++ {
		go worker(w, jobs, results)
	}

	go func(j chan Signal) {
		defer close(j)
		for i := 0; i < numJobs; i++ {
			j <- Signal{
				Status: i,
			}
		}
	}(jobs)

	for i := 0; i < numJobs; i++ {
		fmt.Println(<-results)
	}
}
