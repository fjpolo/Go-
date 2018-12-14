package main

import "fmt"

/* Imports */

/* User structure */
type User struct {
	Id             int
	Name, Location string
}

/* Method */
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s!", u.Name, u.Location)
}

/* Player structure */
type Player struct {
	*User
	GameId int
}

/* Function */
func NewPlayer(id int, name, location string, gameId int) *Player {

	/**/
	return &Player{
		User:   &User{id, name, location},
		GameId: gameId,
	}
}

/*
 * MAIN
 */
func main() {

	/**/
	p := NewPlayer(42, "Matt", "NY", 90404)

	/* Print */
	fmt.Println(p.Greetings())
}
