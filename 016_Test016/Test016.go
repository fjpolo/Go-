package main

import "fmt"

func main() {
	scores := make([]int, 6, 10)
	scores = append(scores, 5)
	fmt.Println(scores)
}
