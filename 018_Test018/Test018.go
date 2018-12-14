package main

import (
	"fmt"
	"math/rand"
	"sort"
)

/**/
func main() {

	/**/
	scores := make([]int, 25)

	/**/
	for i := 0; i < 25; i++ {
		scores[i] = int(rand.Int31n(1000))
	}

	/**/
	sort.Ints(scores)

	/**/
	worst := make([]int, 5)
	copy(worst, scores[:5])

	/**/
	fmt.Println("Original:")
	fmt.Println(scores)
	fmt.Println("Worst cases:")
	fmt.Println(worst)
}
