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
	var length int
	if len(firstArr) > len(secondArr) {
		length = len(secondArr)
	} else {
		length = len(firstArr)
	}
	for i := 0; i < length; i++ {
		switch {
		case i == 0 && firstArr[0] < secondArr[0]:
			newArr[i] = firstArr[i]
			newArr[i+1] = secondArr[i]
		case firstArr[i] <= secondArr[i] && firstArr[i] > secondArr[i-1]:
			newArr[i+i-1+1] = firstArr[i]
			newArr[i+i-1+2] = secondArr[i]
		case firstArr[i] > secondArr[i] && firstArr[i-1] < secondArr[i]:
			newArr[i+i-1+1] = secondArr[i]
			newArr[i+i-1+2] = firstArr[i]
		case firstArr[i] <= secondArr[i] && firstArr[i] < secondArr[i-1]:
			firstArr[i], secondArr[i-1] = secondArr[i-1], firstArr[i]
			newArr[i+i-1] = secondArr[i-1]
			newArr[i+i] = firstArr[i]
			newArr[i+i+1] = secondArr[i]
		case firstArr[i] > secondArr[i] && firstArr[i-1] > secondArr[i]:
			firstArr[i-1], secondArr[i] = secondArr[i], firstArr[i-1]
			newArr[i+i-1] = firstArr[i-1]
			newArr[i+i] = secondArr[i]
			newArr[i+i+1] = firstArr[i]
		}
	}

	if len(firstArr) < len(secondArr) {
		if secondArr[len(firstArr)] <= firstArr[len(firstArr)-1] {
			newArr[len(firstArr)*2-1], newArr[len(firstArr)*2] = secondArr[len(firstArr)], firstArr[(len(firstArr)-1)]
		} else {
			newArr[len(firstArr)*2] = secondArr[len(secondArr)-1]
		}

		for i := firstArr[len(firstArr)-1]; i < len(newArr); i++ {
			newArr[len(firstArr)*2] = secondArr[i]
		}
	}
	return newArr
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Слияние отсортированных массивов.")
	fmt.Println()

	var firstArr = [firstArrSize]int{1, 3, 13, 19}
	secondArr := [secondArrSize]int{4, 5, 11, 12, 16}

	fmt.Printf("Первый массив: %v, второй массив: %v;\nНовый объедененный массив: %v\n", firstArr, secondArr, concatenation(firstArr, secondArr))
}
