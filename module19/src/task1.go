package main

import (
	"fmt"
)

const firstArrSize = 4
const secondArrSize = 5
const newArrSize = firstArrSize + secondArrSize

// Принимает два массива определенной длины и соединяет их в один массив, который возвращает функция.
func concatenation(firstArr [firstArrSize]int, secondArr [secondArrSize]int) (newArr [newArrSize]int) {

	var indexOne, indexTwo, total int

	for indexOne < firstArrSize && indexTwo < secondArrSize {
		if firstArr[indexOne] < secondArr[indexTwo] {
			newArr[total] = firstArr[indexOne]
			indexOne++
		} else {
			newArr[total] = secondArr[indexTwo]
			indexTwo++
		}
		total++
	}

	for indexOne < firstArrSize {
		newArr[total] = firstArr[indexOne]
		indexOne++
		total++
	}

	for indexTwo < secondArrSize {
		newArr[total] = secondArr[indexTwo]
		indexTwo++
		total++
	}
	return
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Слияние отсортированных массивов.")
	fmt.Println()

	var firstArr = [firstArrSize]int{1, 3, 13, 19}
	secondArr := [secondArrSize]int{4, 5, 11, 12, 16}

	fmt.Printf("Первый массив: %v, второй массив: %v;\nНовый объедененный массив: %v\n", firstArr, secondArr, concatenation(firstArr, secondArr))
}
