package main

import (
	"fmt"
)

func calculations(numberFrom *int, numberUpTo *int, result *int) {
	if *numberFrom > *numberUpTo {
		numberFrom, numberUpTo = numberUpTo, numberFrom
	}
	for i := *numberUpTo; i >= *numberFrom; i-- {
		if i%2 == 0 {
			*result = *result + i
		}
	}
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Функция, принимающая аргументы.")
	fmt.Println()

	var numberFrom, numberUpTo, result int

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&numberFrom)
	if err != nil {
		fmt.Print("Неверный ввод! Ошибка: ")
		fmt.Println(err)
	}
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&numberUpTo)
	if err != nil {
		fmt.Print("Неверный ввод! Ошибка: ")
		fmt.Println(err)
	}

	calculations(&numberFrom, &numberUpTo, &result)
	fmt.Println(result)
}
