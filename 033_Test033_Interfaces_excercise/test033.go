/*
Copy your Sqrt function from the earlier exercises (Section 5.4) and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn’t support complex numbers.

Create a new type

type ErrNegativeSqrt float64
and make it an error by giving it a

func (e ErrNegativeSqrt) Error() string
method such that ErrNegativeSqrt(-2).Error() returns

cannot Sqrt negative number: -2.

Note: a call to fmt.Print(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Print(float64(e)). Why?

Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.
*/
package main

import (
	"fmt"
)

/* new type asked*/
type ErrNegativeSqrt float64

/* Methods */
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("0H $H1T! Cannot do Sqrt of a negative number: %g", float64(e))
}

/* Functions */
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((z*z)-x)/(2*z)
	}
	return z, nil
}

/*
 * MAIN
 **/
func main() {
	var input float64

	//Ask for a value
	fmt.Println("Send me a number to squareroot")
	// Scan it
	fmt.Scanf("%f", &input)
	//Print its squareroot
	fmt.Println(Sqrt(input))
}
