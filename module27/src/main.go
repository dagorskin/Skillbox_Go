package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Student struct {
	name  string
	age   int
	grade int
}

// Метод добавления студента в хэш-карту.
func put(students map[string]*Student, student *Student) {
	students[student.name] = student
}

// Метод запрашивает имя студента из хэш-карты.
func get(students map[string]*Student, name string) (*Student, error) {
	if students[name] == nil {
		return nil, errors.New("ошибка: студент не был найден в списке")
	} else {
		return students[name], nil
	}
}

// Информация о студенте.
func (s Student) info() string {
	return s.name + ", " + strconv.Itoa(s.age) + " лет, курс " + strconv.Itoa(s.grade) + ";"
}

func main() {
	students := make(map[string]*Student, 0) // Список студентов.

	fmt.Println("Здравствуйте! Введите имя, возраст и курс студентов через пробел:")

	var counter = 1
	var in = bufio.NewReader(os.Stdin)

	// Ввод данных о студенте.
	for {
		line, err := in.ReadString('\n')

		if err == io.EOF {
			fmt.Println("Ввод закончен! Список студентов:")
			break
		}

		lineFields := strings.Fields(line)

		if len(lineFields) < 3 {
			fmt.Println("Нужно ввести все данные для продолжения! Повторите снова.")
			continue
		}

		studentName := lineFields[0]
		studentAge, errAge := strconv.Atoi(lineFields[1])
		studentGrade, errGrade := strconv.Atoi(lineFields[2])

		if errAge != nil || errGrade != nil {
			fmt.Println("Ошибка при вводе возраста студента и/или его курса! Повторите снова.")
			continue
		}

		student := Student{
			name:  studentName,
			age:   studentAge,
			grade: studentGrade,
		}

		if _, err := get(students, student.name); err != nil {
			put(students, &student)
		} else {
			fmt.Println("Студент с таким именем уже есть в списке! Повторите снова.")
		}
	}

	// Вывод всех студентов из списка.
	for _, value := range students {
		fmt.Printf("%v. %s\n", counter, value.info())
		counter++
	}
}
