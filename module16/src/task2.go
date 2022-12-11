package main

import (
	"fmt"
)

// Один проход сортировки.
func sorting(mainArr [6]int, swap int) ([6]int, int) {
	swap = 0
	for index, _ := range mainArr {
		if index != 0 && mainArr[index] < mainArr[index-1] {
			mainArr[index], mainArr[index-1] = mainArr[index-1], mainArr[index]
			swap++
		}
	}
	return mainArr, swap
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Сортировка пузырьком.")
	fmt.Println()

	var mainArr = [6]int{4, 5, 1, 3, 2, 6}
	var oldArr = mainArr
	var swap = 1

	// Вызов функции пока переменная отвечающая за количество замен не будет равна 0.
	for swap != 0 {
		mainArr, swap = sorting(mainArr, swap)
	}

	fmt.Printf("Оригинальный массив: %v\nОтсортированный массив: %v\n", oldArr, mainArr)
}
