package main

import (
	"bufio"
	"fmt"
	"io"
	storage2 "module28/internal/storage"
	"module28/internal/students"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)
var student students.Student
var position, count = 1, 1

type App struct {
	repository storage2.Storage
}

func NewApp(repository storage2.Storage) *App {
	return &App{repository: repository}
}

func (a *App) run() {
	for {
		if s, ok := a.inputNextStudent(); ok {
			a.storeStudents(s)
		} else {
			a.printStudents()
			break
		}
	}
}

func (a *App) printStudents() {
	studentsPrint := a.repository.Get()
	for _, value := range studentsPrint {
		fmt.Printf("%v. %s\n", position, value.Info())
		position++
	}
}

func (a *App) inputNextStudent() (students.Student, bool) {
	fmt.Println("Введите имя, возраст и курс студента:")
	line, err := in.ReadString('\n')
	fmt.Println()

	if err == io.EOF {
		fmt.Println("Ввод закончен! Список студентов:")
		return student, false
	}

	lineFields := strings.Fields(line)

	if len(lineFields) < 3 {
		fmt.Println("Нужно ввести все данные для продолжения! Повторите снова.")
		return student, false
	}

	studentName := lineFields[0]
	studentAge, errAge := strconv.Atoi(lineFields[1])
	studentGrade, errGrade := strconv.Atoi(lineFields[2])

	if errAge != nil || errGrade != nil {
		fmt.Println("Ошибка при вводе возраста студента и/или его курса! Повторите снова.")
		return student, false
	}

	student = students.MakeStudent(studentName, studentAge, studentGrade)
	count++
	return student, true
}

func (a *App) storeStudents(student students.Student) {
	a.repository.Put(count, student)
}

func main() {
	repository := storage2.NewMemStorage()

	var app = NewApp(repository)

	app.run()

	fmt.Println("...вывод закончен...")
}
