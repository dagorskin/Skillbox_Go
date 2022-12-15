package main

import (
	"fmt"
	"math"
)

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Расчёт по формуле.")
	fmt.Println()

	var (
		S float32
		x int16
		y uint8
		z float32
	)

	S = float32(2) * float32(x) * float32(math.Pow10(int(y))) - float32(3)/z

	fmt.Println("Результат:", S)
}
