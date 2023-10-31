package models

type TeacherDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

var Teacher *TeacherDTO

var TeacherArray []TeacherDTO

func CreateTeacher(id int, firstName string, lastName string, age int) TeacherDTO {
	Teacher = &TeacherDTO{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}

	return *Teacher
}
