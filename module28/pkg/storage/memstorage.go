package storage

import "module28/pkg/students"

type MemStorage struct {
	students map[int]students.Student
}

func NewMemStorage() *MemStorage {
	return &MemStorage{
		students: make(map[int]students.Student),
	}
}

func (ms *MemStorage) Put(count int, student students.Student) {
	ms.students[count] = student
}

func (ms *MemStorage) Get() map[int]students.Student {
	return ms.students
}
