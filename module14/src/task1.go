package main

import (
	"fmt"
)

// Определение чётности.
func parityCheck(number int) bool {
	if number%2 == 0 {
		return true
	} else {
		return false
	}
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Функция, возвращающая результат.")
	fmt.Println()

	// Ввод числа
	var numberInput int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&numberInput)
	if err != nil {
		fmt.Print("Неверный ввод! Ошибка: ")
		panic(err)
	}

	// Определение чётности и вывод результата.
	resultBool := parityCheck(numberInput)
	if resultBool == true {
		resultStr := "чётным"
		fmt.Printf("Число %d является %s числом.\n", numberInput, resultStr)
	} else {
		resultStr := "нечётным"
		fmt.Printf("Число %d является %s числом.\n", numberInput, resultStr)
	}
}
