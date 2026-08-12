package main

import (
	"fmt"
	"sort"
	"sync"
)

// TODO: Implement worker(id int, jobs <-chan int, results chan<- int)
// It should process jobs by multiplying each received job integer by 2,
// and sending the processed value to results.

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// TODO: Spawn 3 worker goroutines using sync.WaitGroup to track completion.

	// TODO: Send 5 jobs (values 1 to 5) and close the jobs channel.

	// TODO: Close the results channel after all workers are done.

	// TODO: Collect and print the results in sorted order to match the expected output.
}
