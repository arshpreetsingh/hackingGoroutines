package main

import (
	"fmt"
	"sync"
	"time"
)

// This time we are using wait-group, Worker will only exit when job will be done, So
// So no need to receive values from result channels. Looks not so good at this tie as well.

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int, errors chan<- error) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)

		if j%2 == 0 {
      fmt.Println(j%2, "jobID",id)
			results <- j * 2
		} else {
      fmt.Println(j%2, "jobID",id)
			errors <- fmt.Errorf("error on job %v", j)
		}
		wg.Done()
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	errors := make(chan error, 100) // Try to make your Channel Smaller,    errors := make(chan error, 1) //and find out what happens!!

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		go worker(w, &wg, jobs, results, errors)
	}

	for j := 1; j <= 900; j++ {
		jobs <- j
		wg.Add(1)
	}
	close(jobs)

	wg.Wait()

	select {
	case err := <-errors:
		fmt.Println("finished with error:", err.Error())
	default:
	}
}
