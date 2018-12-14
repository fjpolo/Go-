package main

import "fmt"

/* Saiyan Struct */
type saiyan struct {
	Name  string
	Power int
}

/* Methods */
func (s *saiyan) super() {
	s.Power += 10000
}

/*
* MAIN
**/
func main() {

	/**/
	goku := &saiyan{"Goku", 9000}
	goku.super()
	/**/
	fmt.Printf("%s: %d", goku.Name, goku.Power)
}
