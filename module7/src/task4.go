package main

import "fmt"

func task4() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Зеркальные билеты.")
	fmt.Println()

	// Инициализация данных разных типов.
	const (
		minLimit int = 100000
		maxLimit int = 999999
	)
	var (
		luckyTicketOld, luckyTicketNew, firstPartTicket, endPartTicket, sumDigitFirstPartTicket,
		sumDigitEndPartTicket, distanceBetweenTicketsOld, distanceBetweenTicketsNew int
	)
	luckyTicketOld = minLimit

	for ticket := minLimit; ticket <= maxLimit; ticket++ {
		luckyTicketNew = ticket
		firstPartTicket = luckyTicketNew / 1000
		endPartTicket = luckyTicketNew % 1000

		for i := 3; i != 0; i-- {
			sumDigitFirstPartTicket = sumDigitFirstPartTicket + firstPartTicket%10
			firstPartTicket /= 10
			sumDigitEndPartTicket = sumDigitEndPartTicket + endPartTicket%10
			endPartTicket /= 10
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
