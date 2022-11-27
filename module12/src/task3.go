package main

import (
	"fmt"
	"os"
)

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

	var fileWrite string
	fmt.Print("Введите текст для записи в файл: ")
	_, _ = fmt.Scan(&fileWrite)

	if _, err = file.WriteString(fileWrite); err != nil {
		panic(err)
	} else {
		fmt.Println("Запись прошла успешно.")
	}
}
