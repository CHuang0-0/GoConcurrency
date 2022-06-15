package main

import (
	"fmt"
	"time"
)

func main() {
	const njobs, nworkers = 5, 3

	//make a channel to send
	jobs := make(chan int, njobs)
	//make a channel to receive
	res := make(chan int, njobs)

	// create 3 goroutines
	for x := 1; x <= nworkers; x++ {
		go workerpool(x, jobs, res)
	}

	// send 5 jobs to the jobs channel which only receives
	for j := 1; j <= njobs; j++ {
		jobs <- j
	}
	//close channel
	close(jobs)
	//send results to the res channel
	for r := 1; r <= njobs; r++ {
		v := <-res
		fmt.Println("results", v)
	}
}

// jobs channel only receives
// res channel only sends
func workerpool(id int, jobs <-chan int, res chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		res <- j * 2 //randomly return some number to res channel
	}
}
