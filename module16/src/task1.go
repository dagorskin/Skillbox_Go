package main

import (
	"fmt"
)

const firstArrSize = 4
const secondArrSize = 5
const newArrSize = firstArrSize + secondArrSize

// Принимает два массива определенной длины и соединяет их в один массив, который возвращает функция.
func concatenation(firstArr [firstArrSize]int, secondArr [secondArrSize]int) [9]int {
	var newArr [newArrSize]int
	for indexOne := 0; indexOne < len(newArr); indexOne++ {
		if indexOne < len(firstArr) {
			newArr[indexOne] = firstArr[indexOne]
		} else {
			for indexTwo := 0; indexTwo != len(secondArr); indexTwo++ {
				newArr[indexOne] = secondArr[indexTwo]
				indexOne++
			}
		}
	}
	return newArr
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Слияние отсортированных массивов.")
	fmt.Println()

	var firstArr = [firstArrSize]int{1, 2, 3, 4}
	secondArr := [secondArrSize]int{5, 6, 7, 8, 9}

	fmt.Printf("Первый массив: %v, второй массив: %v;\nНовый объедененный массив: %v\n", firstArr, secondArr, concatenation(firstArr, secondArr))
}
