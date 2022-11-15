package main

import (
	"fmt"
	"math"
)

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Переполнение.")
	fmt.Println()

	// Инициализация.
	var (
		overflowUnit8Count  int
		overflowUnit16Count int
		varUint8            uint8  = math.MaxUint8
		varUint16           uint16 = math.MaxUint16
	)

	// Подсчет количества переполнений.
	for varUint64 := 0; varUint64 != math.MaxUint32; varUint64++ {
		if varUint8 == 0 {
			overflowUnit8Count++
		}

		if varUint16 == 0 {
			overflowUnit16Count++
		}
		varUint8--
		varUint16--
	}

	// Вывод результатов.
	fmt.Printf("Количество переполнений чисел типа uint8: %d, чисел типа uint16: %d;\n", overflowUnit8Count, overflowUnit16Count)

}
