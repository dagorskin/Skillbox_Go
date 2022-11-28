package main

import (
	"fmt"
	"os"
)

// Создание файла с допуском 0444
func secondOpen(nameFile string) {
	file, err := os.Open(nameFile + ".txt")
	var fileWrite string
	fmt.Print("Введите текст для записи в файл: ")
	_, _ = fmt.Scan(&fileWrite)

	if _, err = file.WriteString(fileWrite); err != nil {
		fmt.Println("Доступ запрещен. Ошибка:", err)

	} else {
		fmt.Println("Запись прошла успешно.")
	}
}

func task3() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 3. Уровни доступа.")
	fmt.Println()

	var nameFile string
	fmt.Print("Введите название создаваемого файла: ")
	_, _ = fmt.Scan(&nameFile)

	file, err := os.Create(nameFile + ".txt")
	if err = os.Chmod(nameFile+".txt", 0444); err != nil {
		fmt.Println("Создать файл не получилось!")
	}
	defer file.Close()

	secondOpen(nameFile)

}
