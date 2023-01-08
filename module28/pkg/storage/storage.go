package storage

import (
	"errors"
	"module28/pkg/students"
)

// Put Метод добавления студента в хэш-карту.
func Put(students map[string]*students.Student, student *students.Student) {
	students[student.Name] = student
}

// Get Метод запрашивает имя студента из хэш-карты.
func Get(students map[string]*students.Student, name string) (*students.Student, error) {
	if students[name] == nil {
		return nil, errors.New("ошибка: студент не был найден в списке")
	} else {
		return students[name], nil
	}
}
