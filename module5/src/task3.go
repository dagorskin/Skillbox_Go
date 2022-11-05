package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 3. Проверить, что есть совпадающие числа.")
	fmt.Println("------------------------------")

	var firstNumber, secondNumber, thirdNumber int

	// Ввод данных.
	fmt.Print("Введите ТРИ числа в одну строку разделяя их пробелом: ")
	fmt.Scanln(&firstNumber, &secondNumber, &thirdNumber)

	if firstNumber == secondNumber && secondNumber == thirdNumber {
		fmt.Println("Все три числа равны!")
	} else if firstNumber == secondNumber {
		fmt.Println("Первое и второе число равны.")
	} else if firstNumber == thirdNumber {
		fmt.Println("Первое и третье число равны.")
	} else if secondNumber == thirdNumber {
		fmt.Println("Второе и третье число равны.")
	} else {
		fmt.Println("Все три числа разные!")
	}
}
