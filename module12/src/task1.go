package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Работа с файлами.")
	fmt.Println()

	// Инициализация.
	var (
		dataInput    string
		dataWrite    string
		numberString uint32
	)

	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("Файл не создан! Ошибка: ")
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for dataInput != "exit" {
		fmt.Printf("Введите какой-либо произвольный текст (exit = выход): ")
		scanner.Scan()
		dataInput = scanner.Text()

		if dataInput == "exit" {
			return
		} else {
			numberString++
		}
		timeNow := time.Now().Format("2006-01-02 15:04:05")
		dataWrite = strconv.Itoa(int(numberString)) + ") " + timeNow + " " + dataInput + "\n"

		if _, err = file.WriteString(dataWrite); err != nil {
			fmt.Print("Ошибка при попытке записи: ")
			fmt.Println(err)
		}
		fmt.Println("Запись прошла успешно...")

	}
}
