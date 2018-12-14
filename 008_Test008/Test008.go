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
	return 9000, false
}

/*** ***/
func main() {

	/**/
	_, exists := power("Goku")

	/**/
	if !exists {
		fmt.Println("ERROR!Goku does not exist!")
	}
}
