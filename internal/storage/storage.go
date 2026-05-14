package storage

import "github.com/tv-anagha/rest-api/internal/types"

//interfaces for storage

type Storage interface {
	CreateStudent(name string, email string, age int) (int, error)
	GetStudentById(id int) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(id int, name string, email string, age int) (int, error)
	DeleteStudentById(id int) (int, error)
}