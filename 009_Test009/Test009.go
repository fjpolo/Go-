package main

import "fmt"

/**/
type saiyan struct {
	Name  string
	Power int
}

/**/
func super(s *saiyan) {
	s.Power += 10000
}

/**/
func main() {

	/**/
	goku := &saiyan{"Goku", 9000}
	super(goku)
	/**/
	fmt.Printf("%s: %d", goku.Name, goku.Power)
}
