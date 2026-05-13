package storage

import "github.com/tv-anagha/rest-api/internal/types"

//interfaces for storage

type Storage interface {
	CreateStudent(name string, email string, age int) (int, error)
	GetStudentById(id int) (types.Student, error)
	// updateStudent(id int, name string, email string, age int) error
	// deleteStudent(id int) error
}