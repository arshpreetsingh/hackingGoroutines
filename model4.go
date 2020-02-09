package main

import (
	"fmt"
	"sync"
	"time"
)

// This time we are using wait-group, Worker will only exit when job will be done, But in
// this model there will be small change,, Which is , we will add some Timeout() case with
// wg.Done, If wg.done() is not returning in that specific case then we have to Exit from that
// specific Go-routines.

func worker(id int, wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	for j := range jobs {
    defer wg.Done()
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second*5)
      //fmt.Println(j%2, "jobID",id)
      results <- j * 2
	}
}

// WaitTimeout does a Wait on a sync.WaitGroup object but with a specified
// timeout. Returns true if the wait completed without timing out, false
// otherwise.

func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration, results chan int) bool {
  go func() {
    wg.Wait()
    close(results)
  }()
  select {
    case <-results:
      return true
    case <-time.After(timeout):
      fmt.Println("exiting because of timeout")
      return false
      }
  }

func main() {
	jobs := make(chan int, 100)
  results := make(chan int, 100)
	// errors := make(chan error, 100) // Try to make your Channel Smaller,    errors := make(chan error, 1) //and find out what happens!!

	var wg sync.WaitGroup
  // start the workers!!
	for w := 1; w <= 4; w++ {
		go worker(w, &wg, jobs, results)
	}
/// Send the Jobs!!
	for j := 1; j <= 20; j++ {
		jobs <- j
		wg.Add(1) // For each go-routine Job, Add one wait Group!!
	}
  // now use the WaitTimeout instead of wg.Wait()
WaitTimeout(&wg, 6 * time.Second,results)
}
