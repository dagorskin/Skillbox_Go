package main

import (
	"fmt"
	"math"
)

func sorting(mainArr [6]int, leftChar, rightChar, count int) {
	for index, char := range mainArr {
		rightChar = char
		if leftChar > rightChar {
			if index == 0 {
				leftChar = char
				continue
			} else {
				mainArr[index], mainArr[index-1] = mainArr[index-1], mainArr[index]
				count++
			}
		}
		leftChar = char
		if index == 0 {
			count = 0
		}
		if count == 0 {
			break
		}
	}
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Сортировка пузырьком.")
	fmt.Println()

	var mainArr = [6]int{4, 5, 1, 3, 2, 6}
	var leftChar, rightChar, count = math.MaxInt, math.MaxInt, 1

	sorting(mainArr, leftChar, rightChar, count)
}
