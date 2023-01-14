package students

import "strconv"

type Student struct {
	Name       string
	Age, Grade int
}

func MakeStudent(name string, age, grade int) Student {
	result := Student{
		Name:  name,
		Age:   age,
		Grade: grade,
	}
	return result
}

// Info Информация о студенте.
func (s Student) Info() string {
	return s.Name + ", возраст " + strconv.Itoa(s.Age) + ", курс " + strconv.Itoa(s.Grade) + ";"
}
