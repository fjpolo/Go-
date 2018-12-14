package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	/* Counter */
	counter = 0
	/* Mutex */
	lock sync.Mutex
)

func main() {

	/* Loop */
	for i := 0; i < 20; i++ {
		/* Call goroutine */
		go incr()
	}

	/* 10mS Delay */
	time.Sleep(time.Millisecond * 10)
}

func incr() {
	/* Lock */
	lock.Lock()
	/* Defer unlock */
	defer lock.Unlock()
	/* Increment counter */
	counter++
	/* Print result */
	fmt.Println(counter)
}
