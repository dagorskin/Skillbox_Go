package main

import (
	"bufio"
	"fmt"
	"io"
	"module28/pkg/storage"
	"module28/pkg/students"
	"os"
	"strconv"
	"strings"
)

func main() {
	studentsMap := make(map[string]*students.Student, 0) // Список студентов.

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

		student := students.Student{
			Name:  studentName,
			Age:   studentAge,
			Grade: studentGrade,
		}

		if _, err := storage.Get(studentsMap, student.Name); err != nil {
			storage.Put(studentsMap, &student)
		} else {
			fmt.Println("Студент с таким именем уже есть в списке! Повторите снова.")
		}
	}

	// Вывод всех студентов из списка.
	for _, value := range studentsMap {
		fmt.Printf("%v. %s\n", counter, value.Info())
		counter++
	}
}
