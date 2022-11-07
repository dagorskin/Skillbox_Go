package main

import (
	"fmt"
	"time"
)

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Зеркальные билеты.")
	fmt.Println("------------------------------")

	// Инициализация данных разных типов.
	var startNumber int = 100000
	var num1, num2, num3, num4, num5, num6, mirrorDigitCounter int

	// Для расчета скорости работы программы.
	start := time.Now()

	// Подсчет количества зеркальных билетов.
	for startNumber <= 999999 {
		num1 = startNumber / 100000
		num2 = startNumber / 10000 % 10
		num3 = startNumber / 1000 % 10
		num4 = startNumber / 100 % 10
		num5 = startNumber / 10 % 10
		num6 = startNumber % 10
		startNumber++

		if num1 == num6 && num2 == num5 && num3 == num4 {
			mirrorDigitCounter++
		}
	}

	// Фиксация скорости работы программы.
	duration := time.Since(start)

	// Вывод результата.
	fmt.Print("Программа насчитала зеркальных чисел: ", mirrorDigitCounter, " за ", duration, ".\n")

}
