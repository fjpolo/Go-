package main

import (
	"fmt"
)

/*
 * User struct
 */
type User struct {
	FirstName, LastName string
}

/* Methods */
func (u *User) Name() string {
	/**/
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

/*
 * Customer struct
 */
type Customer struct {
	Id       int
	FullName string
}

/* Methods */
func (c *Customer) Name() string {
	/**/
	return c.FullName
}

/*
 * Interfaces
 */

/* Namer */
type Namer interface {
	Name() string
}

/* Functions */
func Greet(n Namer) string {
	/**/
	return fmt.Sprintf("Dear %s", n.Name())
}

/*
* MAIN
**/
func main() {
	/* User */
	u := &User{"Matt", "Aimonetti"}
	fmt.Println(Greet(u))
	/* Customer */
	c := &Customer{42, "Francesc"}
	fmt.Println(Greet(c))
}
