package students

import "strconv"

type Student struct {
	Name       string
	Age, Grade int
}

// Info Информация о студенте.
func (s Student) Info() string {
	return s.Name + ", " + strconv.Itoa(s.Age) + " лет, курс " + strconv.Itoa(s.Grade) + ";"
}
