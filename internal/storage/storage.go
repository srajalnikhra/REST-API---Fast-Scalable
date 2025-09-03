package storage

import "github.com/codersgyan/students-api/internal/storage/types"

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(id int64, student types.Student) error
	DeleteStudent(id int64) error
}