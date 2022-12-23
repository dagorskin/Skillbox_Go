package main

import (
	"fmt"
)

func sort(unsortedArray [10]int) [10]int {
	var digit, count int
	for index := 2; index < len(unsortedArray); index++ {
		digit = unsortedArray[index]
		count = index - 1
		for count >= 0 && unsortedArray[count] > digit {
			unsortedArray[count+1] = unsortedArray[count]
			count -= 1
		}
		unsortedArray[count+1] = digit
	}
	return unsortedArray
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Сортировка вставками.")
	fmt.Println()

	theArray := [10]int{4, 7, 1, 9, 2, 3, 6, 5, 10, 8}

	fmt.Println("Неотсортированный массив:\n", theArray)
	fmt.Println("Отсортированный массив:\n", sort(theArray))
}
