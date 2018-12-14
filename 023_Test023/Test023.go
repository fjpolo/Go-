package main

import (
	"fmt"
	"hospital"
)

/* Main */
func main() {

	/*  */
	myPacient := hospital.Patient{}
	myPacient.Name = "Juan"
	myPacient.Surname = "Perez"
	myPacient.Age = 52
	myPacient.Height = 1.78
	myPacient.Weight = 78.2
	myPacient.InsuranceNr = 123

	/**/
	//fmt.Println(myPacient)
	fmt.Println(myPacient.GetInfo())
	//myPacient.GetInfo()
}
