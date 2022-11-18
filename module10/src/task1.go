package main

import (
	"fmt"
	"math"
)

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Разложение ex в ряд Тейлора.")
	fmt.Println()

	// Инициализация.
	var (
		i                      uint64
		numberX                int
		accuracy               int
		resultTaylorSeries     float64 = 1.0
		resultTaylorSeriesPrev float64
		fact                   uint64 = 1 // Факториал.
	)

	// Ввод данных.
	fmt.Print("Введите число Х: ")
	_, _ = fmt.Scan(&numberX)
	fmt.Print("Введите точность: ")
	_, _ = fmt.Scan(&accuracy)
	epsilon := 1 / math.Pow10(accuracy)

	// Расчет #1
	for math.Abs(resultTaylorSeries-resultTaylorSeriesPrev) > epsilon {
		i++
		fact *= i
		resultTaylorSeriesPrev = resultTaylorSeries
		resultTaylorSeries += math.Pow(float64(numberX), float64(i)) / float64(fact)
	}

	// Расчет #2
	/*for i := 1; i <= accuracy; i++ {
		fact *= float64(i)
		resultTaylorSeries += math.Pow(float64(numberX), float64(i)) / fact
		if math.Abs(resultTaylorSeries-resultTaylorSeriesPrev) < epsilon {
			break
		}
		resultTaylorSeriesPrev = resultTaylorSeries
	}*/

	// Вывод результата.
	fmt.Println("Экспонента =", resultTaylorSeries)
}
