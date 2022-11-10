package main

import (
	"fmt"
	"time"
)

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Зеркальные билеты.")
	fmt.Println()

	// Инициализация данных разных типов.
	const startNumberTicket = 100000
	var mirrorDigitCounter, firstPartTicket, lastChar, reversedSecondPartTicket int

	// Для расчета скорости работы программы.
	start := time.Now()

	// Подсчет количества зеркальных билетов.
	for ticket := startNumberTicket; ticket <= 999999; ticket++ {
		firstPartTicket = ticket / 1000
		reversedSecondPartTicket = 0
		ticketNew := ticket
		for i := 3; i != 0; i-- {
			lastChar = ticketNew % 10
			reversedSecondPartTicket = reversedSecondPartTicket*10 + lastChar
			ticketNew = ticketNew / 10
		}

		if firstPartTicket == reversedSecondPartTicket {
			mirrorDigitCounter++
		}
	}

	// Фиксация скорости работы программы.
	duration := time.Since(start)

	// Вывод результата.
	fmt.Print("Программа насчитала зеркальных чисел: ", mirrorDigitCounter, " за ", duration, ".\n")

}
