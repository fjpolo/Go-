package main

import "fmt"

/* Structure */
type User struct{
  FirstName,
  LastName string
}

/* Method */
func (u *User) Name() string{
  return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

/* Interface */
type Namer interface{
  Name() string
}

/* Method */
func Greet(n Namer) string{
  return fmt.Sprintf("Dear %s", n.Name())
}

/* MAIN */
func main(){

  /**/
  u := &User{"Juan", "Perez"}

  /**/
  fmt.Println(Greet(u))
}
