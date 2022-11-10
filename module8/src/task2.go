package main

import "fmt"

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Дни недели.")
	fmt.Println()

	// Ввод данных.
	var weekday string
	fmt.Print("Введите день недели (пн, вт, ср, чт, пт): ")
	_, _ = fmt.Scan(&weekday)

	// Выполнение Switch.
	switch weekday {
	case "Пн", "пн":
		fmt.Print("понедельник, ")
		fallthrough
	case "Вт", "вт":
		fmt.Print("вторник, ")
		fallthrough
	case "Ср", "ср":
		fmt.Print("среда, ")
		fallthrough
	case "Чт", "чт":
		fmt.Print("четверг, ")
		fallthrough
	case "Пт", "пт":
		fmt.Println("пятница;")
	default:
		fmt.Println("Неверный ввод! Повторите.")
		task2()
	}
}
