package main

import "fmt"

/* Saiyan Struct */
type saiyan struct {
	Name   string
	Power  int
	Father *saiyan
}

/* Methods */
func (s *saiyan) super() {
	s.Power += 10000
}

/*
* MAIN
**/
func main() {

	/* Goku */
	goku := new(saiyan)
	goku.Name = "Goku"
	goku.Power = 9000
	/* Gohan */
	gohan := new(saiyan)
	gohan.Name = "Gohan"
	gohan.Power = 2000
	gohan.Father = goku
	/* Super saiyan */
	goku.super()
	/**/
	fmt.Printf("%s's power: %d\n", goku.Name, goku.Power)
	fmt.Printf("%s's father: %s\n", gohan.Name, gohan.Father.Name)

}
