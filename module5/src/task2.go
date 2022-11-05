package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Проверить, что одно из чисел — положительное.")
	fmt.Println("------------------------------")

	var firstNumber, secondNumber, thirdNumber int

	// Ввод данных.
	fmt.Print("Введите ТРИ числа в одну строку разделяя их пробелом: ")
	fmt.Scanln(&firstNumber, &secondNumber, &thirdNumber)

	// Решение №1 Проще, как по мне.
	/*var positiveCount int

	if firstNumber > 0 {
		positiveCount++
	}
	if secondNumber > 0 {
		positiveCount++
	}
	if thirdNumber > 0 {
		positiveCount++
	}

	fmt.Print("Количество положительных чисел: ", positiveCount, ".")*/

	// Решение №2 По теме урока.
	if firstNumber > 0 && secondNumber > 0 && thirdNumber > 0 {
		fmt.Print("Все три числа положительные.")
	} else if firstNumber > 0 && secondNumber > 0 || firstNumber > 0 && thirdNumber > 0 || secondNumber > 0 && thirdNumber > 0 {
		fmt.Print("Два числа являются положительными.")
	} else if firstNumber > 0 && secondNumber <= 0 && thirdNumber <= 0 || firstNumber <= 0 && secondNumber > 0 && thirdNumber <= 0 || firstNumber <= 0 && secondNumber <= 0 && thirdNumber > 0 {
		fmt.Print("Всего лишь одно число является положительным.")
	} else {
		fmt.Println("Неверный ввод!")
	}
}
