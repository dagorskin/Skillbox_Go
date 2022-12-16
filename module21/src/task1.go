package main

import (
	"fmt"
	"math"
)

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Расчёт по формуле.")
	fmt.Println()

	var numberOne, numberTwo ,numberThree int
	var S float32

	fmt.Print("Введите три значения через пробел (= x, z ,y): ")
	_, _ = fmt.Scan(&numberOne, &numberTwo, &numberThree)

	x := int16(numberOne)
	y := uint8(numberTwo)
	z := float32(numberThree)

	S = float32(2) * float32(x) * float32(math.Pow10(int(y))) - float32(3)/z

	fmt.Println("Результат:", S)
}
