package main

import (
	"fmt"
	"sync"
	"time"
)

// simulate expensive job
func worker_pool_fn(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func workeWaitGroup(id int) {
	fmt.Printf("Worker in WG %d starting...\n\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker in WG %d done!\n\n", id)
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// inititalize 3 blocked workers
	// once a job is sent, one worker will unblock itself and start
	// once another is sent, next takes it
	for w := 1; w <= 3; w++ {
		go worker_pool_fn(w, jobs, results)
	}

	// send job
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// once a worker tries to receive but its closed, for loop ends and worker finishes
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		fmt.Println("TRY FETCH RESULTS (", a, ")")
		<-results
	}

	// WAIT GROUPS

	fmt.Println("\n\nWAIT GROUPS:\n\n")
	time.Sleep(time.Second)

	// define a wait group with sync.WaitGroup

	var wg sync.WaitGroup

	// IF PASSED TO A FUNCT AS A VARIABLE --> &wg always must be a pointer

	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			workeWaitGroup(i)
		})
	}

	// block until all tasks are done
	wg.Wait()

	// NOTICE THAT this method cannot propagate worker errors, for a more advanced alternative CHECK errgroup package (https://pkg.go.dev/golang.org/x/sync/errgroup)

}
