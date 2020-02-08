package main

import (
  "fmt"
  "time"
  )


// This time in the second model we will also use one more Channel, It is called Error's Channe,
// So our worker  will not Crash and we will also get if there is an error in the system.

func worker(id int, jobsChnl <-chan int, resultsChan chan<- int, errorChan chan<- error ){  // Don't forget to see small difference between Sender and Receiver Channles. :) Quite hacky it is

  for j:=range jobsChnl {
    fmt.Println(id, "*** This is ID of the JOB", j, "*** and this is job")
    time.Sleep(time.Second*4)
    if j%2==0{
      resultsChan <- j * 2
    } else{
      errorChan<-fmt.Errorf("error on job %v", j)
    }

  }
}

func main() {

  jobsc:=make(chan int,100) // why we are using buffered cannels here?
  resultsc:=make(chan int,100)
  errors :=make(chan error,100)

    fmt.Println("hurray!! world!!!")
        // Now , Make sure How many number of workers you want to Start here:
        // Now we have just started three workers
        for w := 1; w <= 10; w++ {
        		go worker(w, jobsc, resultsc,errors)
        	}
          /// Now workers are there, Those need jobs, we can also call it Job's Pool.
          for j:=0;j<900;j++ {
            jobsc<-j /// Send all jobs to calculate to jobsChnl
          }
          close(jobsc) // Close this channel when all jobs has been Sent!!

  	select {
  	case err := <-errors:
  		fmt.Println("finished with error:", err.Error())
  	default:
      <-resultsc
  	}
  }
