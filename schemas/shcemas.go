package schemas

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model // Includes the struct Model in Student
	Name string  `json:"name"`
	CPF int      `json:"cpf"`
	Email string `json:"email"`
	Age int      `json:"age"`
	Active bool  `json:"active"`
}
