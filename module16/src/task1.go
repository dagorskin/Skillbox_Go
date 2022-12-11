package main

import (
	"fmt"
)

// Принимает два массива определенной длины и соединяет их в один массив, который возвращает функция.
func concatenation(firstArr [4]int, secondArr [5]int) [9]int {
	newArr := [9]int{firstArr[0], firstArr[1], firstArr[2], firstArr[3],
		secondArr[0], secondArr[1], secondArr[2], secondArr[3], secondArr[4]}
	return newArr
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Слияние отсортированных массивов.")
	fmt.Println()

	var firstArr = [4]int{1, 2, 3, 4}
	secondArr := [5]int{5, 6, 7, 8, 9}

	fmt.Printf("Первый массив: %v, второй массив: %v;\nНовый объедененный массив: %v\n", firstArr, secondArr, concatenation(firstArr, secondArr))
}
