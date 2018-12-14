package main

import "fmt"

/* Person structure */
type Person struct {
	Name string
}

// Methods
//
func (p *Person) introduce() {

	/**/
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

/* Saiyan structure */
type Saiyan struct {
	*Person
	Power  int
	Father *Saiyan
}

// Methods
func (s *Saiyan) super() {
	s.Power += 10000
}

/*
* MAIN
**/
func main() {

	/* Goku */
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9000,
		Father: nil,
	}

	/**/
	goku.introduce()

}
