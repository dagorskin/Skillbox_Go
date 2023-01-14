package storage

import "module28/pkg/students"

type Storage interface {
	Put(int, students.Student)
	Get() map[int]students.Student
}
