package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* Declare Worker struct */
type Worker struct {
	id int
}

/* Process */
func (w Worker) process(c chan int) {

	/* Infinite loop as in any RTOS */
	for {

		/* Read channel */
		data := <-c
		/* Print channel data */
		fmt.Printf("Worker %d got %d\n", w.id, data)
	}
}

/* MAIN function */
func main() {

	/* Create a channel */
	c := make(chan int)

	/* Create processes - goroutines */
	for i := 0; i < 5; i++ {

		/* Assign a Worker ID */
		worker := &Worker{id: i}

		/*
		*  Once the 5 goroutines have been created,
		* user has no control over the routines.
		 */

		/* Start process for a given ID */
		go worker.process(c)
	}

	/* Infinite loop */
	for {

		/* Send random number to channel */
		c <- (rand.Int())

		/* Sleep for a while */
		time.Sleep(time.Millisecond * 500)
	}
}
