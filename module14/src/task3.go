package main

import (
	"fmt"
)

// Умножение.
func multiplication(numberOne, numberTwo int) (result int) {
	result = numberOne * numberTwo
	return
}

// Сложение.
func summation(numberOne, numberTwo int) (result int)  {
	result = numberOne + numberTwo
	return
}

func task3() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 3. Именованные возвращаемые значения.")
	fmt.Println()

	var firstNumberInput int
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&firstNumberInput)
	if err != nil {
		fmt.Print("Ошибка ввода :")
		fmt.Println(err)
	}

	var secondNumberInput int
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&secondNumberInput)
	if err != nil {
		fmt.Print("Ошибка ввода :")
		fmt.Println(err)
	}

	fmt.Printf("Результаты: умножение - %d, сложение - %d.\n",
		multiplication(firstNumberInput, secondNumberInput), summation(firstNumberInput, secondNumberInput))

}
