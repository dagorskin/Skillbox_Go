package main

import (
	"fmt"
	"math"
)

// Функция-шаблон вывода результата.
func printResult(multiplicationInputNumbers int, resultType string) {
	_, _ = fmt.Printf("Для произведения этих чисел (%d) подойдет тип <%s>.\n", multiplicationInputNumbers, resultType)
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Минимальный тип данных.")
	fmt.Println()

	// Инициализация.
	var (
		inputNumberOne             int
		inputNumberTwo             int
		multiplicationInputNumbers int
	)

	// Ввод данных.
	fmt.Print("Введите первое число: ")
	_, _ = fmt.Scan(&inputNumberOne)
	fmt.Print("Введите второе число: ")
	_, _ = fmt.Scan(&inputNumberTwo)

	multiplicationInputNumbers = inputNumberOne * inputNumberTwo

	// Поиск нужного типа.
	switch {
	case 0 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxUint8:
		printResult(multiplicationInputNumbers, "uint8")
	case 0 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxUint16:
		printResult(multiplicationInputNumbers, "uint16")
	case 0 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxUint32:
		printResult(multiplicationInputNumbers, "uint32")
	case 0 <= multiplicationInputNumbers:
		printResult(multiplicationInputNumbers, "uint64")
	case math.MinInt8 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxInt8:
		printResult(multiplicationInputNumbers, "int8")
	case math.MinInt16 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxInt16:
		printResult(multiplicationInputNumbers, "int16")
	case math.MinInt32 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxInt32:
		printResult(multiplicationInputNumbers, "int32")
	case math.MinInt64 <= multiplicationInputNumbers && multiplicationInputNumbers < math.MaxInt64:
		printResult(multiplicationInputNumbers, "int64")
	}
}
