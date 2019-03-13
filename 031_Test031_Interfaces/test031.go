package main

import (
	"fmt"
	"os"
)

/*
 * Interfaces
 */

/* Reader */
type Reader interface {
	/**/
	Read(b []byte) (n int, err error)
}

/* Writer */
type Writer interface {
	/**/
	Write(b []byte) (n int, err error)
}

/* ReadWriter */
type ReadWriter interface {
	/**/
	Reader
	Writer
}

/*
 * MAIN
 **/
func main() {
	var w Writer
	/**/
	// os.Stdout implements Writer
	w = os.Stdout
	fmt.Fprintf(w, "hello, writer\n")
}
