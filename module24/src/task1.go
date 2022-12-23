package main

import "fmt"

func justFunc(inputArr ...int) (evenNumberedArr [5]int, unevenNumberArr [5]int) {
	index1, index2 := 0, 0
	for index, number := range inputArr {
		if number%2 == 0 {
			evenNumberedArr[index1] = inputArr[index]
			index1++
		} else {
			unevenNumberArr[index2] = inputArr[index]
			index2++
		}
	}
	return
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Чётные и нечётные.")
	fmt.Println()

	evenNumberedArr, unevenNumberArr := justFunc(2, 5, 11, 8, 7, 12, 20, 15, 22, 21)
	fmt.Printf("Четный массив:\n%v\nНечетный массив:\n%v\n", evenNumberedArr, unevenNumberArr)
}
