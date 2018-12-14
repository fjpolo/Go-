package main

import "fmt"

/* Saiyan Struct */
type saiyan struct {
	Name  string
	Power int
}

/* Methods */
//
func NewSaiyan(name string, power int) *saiyan {
	return &saiyan{
		Name:  name,
		Power: power,
	}
}

//
func (s *saiyan) super() {
	s.Power += 10000
}

/*
 * MAIN
 *
 */
func main() {

	/* New saiyan */
	goku := NewSaiyan("Goku", 9000)
	/* Super saiyan */
	goku.super()
	/* Print it */
	fmt.Printf("%s: %d", goku.Name, goku.Power)
}
