package main

import (
	"fmt"
	"math"
)

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Проблемы округления процентов.")
	fmt.Println()

	// Инициализация.
	var (
		debitClient            float32 // Счет клиента.
		debitBank              float32 // Счет банка.
		capitalization         float32 // Процент капитализации.
		capitalizationPerMonth float32 // Сумма капитализации (заработка) за месяц.
		afterAdjustment        float32 // После округления.
		years                  uint8   // Срок действия вклада.
	)

	// Ввод данных.
	fmt.Println("Вас приветствует мастер вкладов банка <MoneyBox!>")
	fmt.Print("Введите сумму вклада: ")
	_, _ = fmt.Scan(&debitClient)
	fmt.Print("Введите ежемесячный процент капитализации: ")
	_, _ = fmt.Scan(&capitalization)
	fmt.Print("Введите срок вклада в годах: ")
	_, _ = fmt.Scan(&years)

	// Расчет.
	for month := years * 12; month != 0; month-- {
		capitalizationPerMonth = debitClient * capitalization / 100
		debitClient += capitalizationPerMonth
		afterAdjustment = debitClient * 100
		afterAdjustment = float32(math.Trunc(float64(afterAdjustment))) / 100
		debitBank += debitClient - afterAdjustment
		debitClient = afterAdjustment
	}

	// Вывод результата.
	fmt.Printf("Вы заработали за весь срок: %g;\nБанк присвоил себе: %g", debitClient, debitBank)
}
