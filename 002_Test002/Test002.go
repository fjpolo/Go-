package main

import (
	"fmt"
	"os"
)

func main() {

	/**/
	if len(os.Args) != 2 {
		fmt.Println("It's over", os.Args[1])
	}
}
