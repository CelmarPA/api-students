package schemas

import (
	"gorm.io/gorm"
	"time"
)

type Student struct {
	gorm.Model // Includes the struct Model in Student
	Name string  `json:"name"`
	CPF int      `json:"cpf"`
	Email string `json:"email"`
	Age int      `json:"age"`
	Active bool  `json:"active"`
}

type StudentResponse struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
	// DeletedAt time.Time `json:"deleteAt"`
	Name string  `json:"name"`
	CPF int      `json:"cpf"`
	Email string `json:"email"`
	Age int      `json:"age"`
	Active bool  `json:"active"`
}

func NewResponse(students []Student) []StudentResponse {
	studentsResponse := []StudentResponse{}

	for _, student := range students {
		studentResponse := StudentResponse{
			ID: int(student.ID),
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
			Name: student.Name,
			CPF: student.CPF,
			Email: student.Email,
			Age: student.Age,
			Active: student.Active,
		}
		studentsResponse = append(studentsResponse, studentResponse)
	}

	return studentsResponse
}