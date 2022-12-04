package main

import (
	"fmt"
)

func rotation(varFirst *int, varSecond *int) {
	*varFirst, *varSecond = *varSecond, *varFirst
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Функция, принимающая значение по ссылке.")
	fmt.Println()

	var varFirst, varSecond int

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&varFirst)
	if err != nil {
		fmt.Print("Неверный ввод! Ошибка: ")
		fmt.Println(err)
	}
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&varSecond)
	if err != nil {
		fmt.Print("Неверный ввод! Ошибка: ")
		fmt.Println(err)
	}

	fmt.Printf("Первое число = %d, второе число = %d.\n", varFirst, varSecond)
	fmt.Println("Меняем местами...")

	rotation(&varFirst, &varSecond)

	fmt.Printf("Первое число = %d, второе число = %d.\n", varFirst, varSecond)
}
