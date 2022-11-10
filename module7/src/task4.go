package main

import "fmt"

func task4() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 4 (по желанию). Счастливые билеты.")
	fmt.Println()

	// Инициализация данных разных типов.
	const (
		minLimit int = 100000
		maxLimit int = 999999
	)
	var (
		luckyTicketNew, sumDigitFirstPartTicket, sumDigitEndPartTicket,
		distanceBetweenTicketsOld, distanceBetweenTicketsNew int
	)
	luckyTicketOld := minLimit

	for luckyTicketNew = minLimit; luckyTicketNew <= maxLimit; luckyTicketNew++ {
		tempLuckyTicketNew := luckyTicketNew

		for i := 3; i != 0; i-- {
			sumDigitFirstPartTicket += tempLuckyTicketNew % 10
			sumDigitEndPartTicket += tempLuckyTicketNew / 1000 % 10
			tempLuckyTicketNew /= 10
		}

		// Определение большего номера билета и расстояния
		if sumDigitFirstPartTicket == sumDigitEndPartTicket {
			if luckyTicketNew > luckyTicketOld {
				distanceBetweenTicketsNew = luckyTicketNew - luckyTicketOld
				luckyTicketOld = luckyTicketNew
				if distanceBetweenTicketsOld < distanceBetweenTicketsNew {
					distanceBetweenTicketsOld = distanceBetweenTicketsNew
				}
			}
		} else {
			sumDigitFirstPartTicket = 0
			sumDigitEndPartTicket = 0
		}
	}

	fmt.Print("Билет с максимальным счастливым числом - ", luckyTicketOld,
		"\nи максимальное количество билетов - ", distanceBetweenTicketsOld, ".\n")

}
