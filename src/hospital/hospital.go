package hospital

/* Imported packages*/
import . "fmt"

/* Variables */

/* Structures */

//
type Person struct {
	Name,
	Surname string
	Age int
	Height,
	Weight float32
}

//
type Patient struct {
	Person
	InsuranceNr int32
}

/* Functions */

/* Methods */
func (patient *Patient) GetInfo() string {

	/**/
	return Sprintf("Name: %s Surname: %s Age: %d Height: %f Weight: %f Insurance: %d",
		patient.Name, patient.Surname,
		patient.Age, patient.Height,
		patient.Weight, patient.InsuranceNr)
}
