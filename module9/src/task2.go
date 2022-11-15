package main

import (
	"fmt"
	"math"
)

// Функции-шаблоны вывода результата.
func printResultU(multiplicationInputNumbers uint64, resultType string) {
	_, _ = fmt.Printf("Для произведения этих чисел (%d) подойдет тип <%s>.\n", multiplicationInputNumbers, resultType)
}
func printResult(multiplicationInputNumbers int64, resultType string) {
	_, _ = fmt.Printf("Для произведения этих чисел (%d) подойдет тип <%s>.\n", multiplicationInputNumbers, resultType)
}

// Основная функция.
func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Минимальный тип данных.")
	fmt.Println()

	// Инициализация.
	var inputNumberOne, inputNumberTwo int16

	// Ввод данных.
	fmt.Print("Введите первое число: ")
	_, _ = fmt.Scan(&inputNumberOne)
	fmt.Print("Введите второе число: ")
	_, _ = fmt.Scan(&inputNumberTwo)

	// Нахождение типа и вывод результата.
	if inputNumberOne >= 0 && inputNumberTwo >= 0 {
		multiplicationInputNumbers := uint64(inputNumberOne) * uint64(inputNumberTwo)
		switch {
		case 0 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxUint8:
			printResultU(multiplicationInputNumbers, "uint8")
		case 0 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxUint16:
			printResultU(multiplicationInputNumbers, "uint16")
		case 0 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxUint32:
			printResultU(multiplicationInputNumbers, "uint32")
		case 0 <= multiplicationInputNumbers:
			printResultU(multiplicationInputNumbers, "uint64")
		}
	} else {
		multiplicationInputNumbers := int64(inputNumberOne) * int64(inputNumberTwo)
		switch {
		case math.MinInt8 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxInt8:
			printResult(multiplicationInputNumbers, "int8")
		case math.MinInt16 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxInt16:
			printResult(multiplicationInputNumbers, "int16")
		case math.MinInt32 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxInt32:
			printResult(multiplicationInputNumbers, "int32")
		default:
			printResult(multiplicationInputNumbers, "int64")
		}
	}
}
