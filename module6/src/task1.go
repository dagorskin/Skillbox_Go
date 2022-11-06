package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Написание простого цикла.")
	fmt.Println("------------------------------")

	var startNum, endNumber int

	// Ввод данных.
	fmt.Print("Введите число ОТ которого будет вестись отчет: ")
	fmt.Scan(&startNum)
	fmt.Print("Введите число ДО которого будет вестись отчет: ")
	fmt.Scan(&endNumber)

	// Ветвление в зависимости от большего числа, которое ввел пользователь.
	if startNum < endNumber {
		for startNum != endNumber {
			fmt.Println(startNum)
			startNum++
		}
	} else if startNum > endNumber {
		for startNum != endNumber {
			fmt.Println(startNum)
			startNum--
		}
	}
}
