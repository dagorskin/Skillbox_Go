package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Генерация чисел с помощью отсрочки выполнения кода.
func generatingPoints() int {

	t := time.NewTimer(time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	coordinate := -100 + rand.Intn(100-(-100)+1)
	<-t.C
	return coordinate
}

// Изменение точек с помощью формул.
func transformation(pointXOld, pointYOld int) (int, int) {
	pointXNew := 2*pointXOld + 10
	pointYNew := -3*pointYOld - 5
	return pointXNew, pointYNew
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Функция, возвращающая несколько значений.")
	fmt.Println()

	var (
		onePointX   = generatingPoints()
		onePointY   = generatingPoints()
		twoPointX   = generatingPoints()
		twoPointY   = generatingPoints()
		threePointX = generatingPoints()
		threePointY = generatingPoints()
	)

	fmt.Println("Начальные координаты:")
	fmt.Printf("Первая точка: %d, %d;\nВторая точка: %d, %d;\nТретья точка: %d, %d;\n",
		onePointX, onePointY, twoPointX, twoPointY, threePointX, threePointY)

	onePointX, onePointY = transformation(onePointX, onePointY)
	twoPointX, twoPointY = transformation(twoPointX, twoPointY)
	threePointX, threePointY = transformation(threePointX, threePointY)

	fmt.Println("\nИзмененные координаты координаты:")
	fmt.Printf("Первая точка: %d, %d;\nВторая точка: %d, %d;\nТретья точка: %d, %d;\n",
		onePointX, onePointY, twoPointX, twoPointY, threePointX, threePointY)
}
