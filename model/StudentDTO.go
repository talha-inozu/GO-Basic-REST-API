package models

type StudentDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func CreateStudent(id int, firstName string, lastName string, age int) StudentDTO {
	return StudentDTO{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}
