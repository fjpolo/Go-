package main

import "fmt"

/**/
func log(message string) {

}

/**/
func add(a, b int) int {

	/**/
	return a + b
}

/**/
func power(name string) (int, bool) {

	/**/
	return 9000, true
}

/*** ***/
func main() {

	/**/
	value, exists := power("Goku")

	/**/
	if !exists {
		fmt.Println("ERROR!Goku does not exist!")
	} else {
		fmt.Printf("Power of %s: %d\n", "Goku", value)
	}
}
