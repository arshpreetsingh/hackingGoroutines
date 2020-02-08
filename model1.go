package main

import (
  "fmt"
  "time"
  )


// It has two channels one is used to receive jobs/data-for computation
// Other_channel is used to send back Calculated results.
// We willl also use ID-jobID to identyfy what is happening under the hood.


func worker(id int, jobsChnl <-chan int, resultsChan chan<- int ){  // Don't forget to see small difference between Sender and Receiver Channles. :) Quite hacky it is

  for j:=range jobsChnl {
    fmt.Println(id, "*** This is ID of the JOB", j, "*** and this is job")
    time.Sleep(time.Second*30)
    resultsChan <- j * 2
  }
}


func main() {

  jobsc:=make(chan int,100) // why we are using buffered cannels here?
  resultsc:=make(chan int,100)

    fmt.Println("hurray!! world!!!")
        // Now , Make sure How many number of workers you want to Start here:
        // Now we have just started three workers
        for w := 1; w <= 3; w++ {
        		go worker(w, jobsc, resultsc)
        	}
          /// Now workers are there, Those need jobs, we can also call it Job's Pool.
          for j:=1;j<10;j++ {
            jobsc<-j /// Send all jobs to calculate to jobsChnl
          }
          close(jobsc) // Close this channel when all jobs has been Sent!!
          for r:=1;r<10;r++{
            <-resultsc
          }
  }
