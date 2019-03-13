package main

import (
	"fmt"
	"time"
)

/*
 * MyError struct
 */
type MyError struct {
	When time.Time
	What string
}

/* Methods */
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

/*
 * Functions
 */
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

/*
 * MAIN
 **/
func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
