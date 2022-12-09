package main

import (
	"fmt"
)

func inputNumbers() [100]int {
	var countNumbers int
	fmt.Print("Введите количество целых чисел (max 100): ")
	_, err := fmt.Scan(&countNumbers)
	if err != nil {
		fmt.Print("Ошибка ввода: ")
		fmt.Println(err)
	}

	var numbersArr [100]int
	fmt.Printf("Введите %v чисел по порядку:\n", countNumbers)
	for i := 0; i < countNumbers; i++ {
		_, err = fmt.Scan(&numbersArr[i])
		if err != nil {
			fmt.Print("Ошибка ввода: ")
			fmt.Println(err)
		}
	}
	return numbersArr
}

func counting(numbersArr [100]int) (int, int) {
	var (
		countEven int
		countOdd  int
	)
	for _, num := range numbersArr {
		if num == 0 {
			continue
		} else if num%2 == 0 {
			countEven++
		} else {
			countOdd++
		}
	}
	return countEven, countOdd
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Подсчёт чётных и нечётных чисел в массиве.")
	fmt.Println()

	countEven, countOdd := counting(inputNumbers())
	fmt.Printf("Четных чисел: %v\nНечетных чисел: %v\n", countEven, countOdd)
}
