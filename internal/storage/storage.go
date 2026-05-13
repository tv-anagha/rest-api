package storage

//interfaces for storage

type Storage interface {
	CreateStudent(name string, email string, age int) (int, error)
	// getStudent(id int) (*Student, error)
	// updateStudent(id int, name string, email string, age int) error
	// deleteStudent(id int) error
}