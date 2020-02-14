package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// what is structure of task...

type Job struct {
	jobID    int
	randomNo int
}

// What will be produced after processing on Task!!
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 20)       // type of channel which will accept Job{}
var results = make(chan Result, 20) //type of channel to accept Result{}

// will return Doubles of the same number!!!
func Doubles(number int) int {
	ans := number * number
	return ans
}

// Worker to run operation on each Job and return Result!!
func Worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, Doubles(job.randomNo)}
		time.Sleep(time.Second * 1)
		results <- output
	}
	wg.Done() // Only return when Done...! Means when "results" channel will receive Value!!
}

// Now We have to Create Dispatcher which will start all the workers , depends on How many workers we want to start
func CreateWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go Worker(&wg) // Run the worker to as Go Routine!!
	}
	wg.Wait()
	close(results)
}

func SubmitJobs(noOfJobs int) { // Send the Jobs to Channel so able to Process those, Think of it liKe Worker Pool!!
	for i := 0; i < noOfJobs; i++ {
		randno := rand.Intn(999)
		job := Job{i, randno}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Println("Here we are getting results of all the Submitted Jobs!!")
		fmt.Println(result)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	// First We have to Admit Number of Jobs to Pool.
	numberofjobs := 100
	go SubmitJobs(numberofjobs)
	// Now start workers
	done := make(chan bool)
	go result(done)

	numberofworkers := 10
	go CreateWorkerPool(numberofworkers)
	<-done
	// Jobs submitted and Worker is also started! Now wait for results and see magic

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Total Time Taken is::", diff)
}
