package main

import (
	"fmt"
)

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 6. Счастливый билет.")
	fmt.Println("------------------------------")

	// Инициализация переменных.
	var ticketNumber, firstChar, secondChar, thirdChar, fourthChar int

	// Ввод данных.
	fmt.Print("Введите четырехзначный номер билета, чтобы определить его ценность: ")
	fmt.Scan(&ticketNumber)

	// Разбиение номера билета на цифры.
	firstChar = ticketNumber / 1000
	secondChar = ticketNumber / 100 % 10
	thirdChar = ticketNumber / 10 % 10
	fourthChar = ticketNumber % 10

	// Определение типа номера билета.
	if firstChar == fourthChar && secondChar == thirdChar {
		fmt.Println("Этот билет зеркальный!")
	} else if firstChar+secondChar == thirdChar+fourthChar {
		fmt.Println("Этот билет счастливый! Поздравляем!")
	} else {
		fmt.Println("Этот билет обычный.")
	}
}
