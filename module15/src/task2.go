package main

import (
	"fmt"
)

func inputNumbersNew() ([]int, int) {
	var countNumbers int
	fmt.Print("Введите количество целых чисел (max 100): ")
	_, err := fmt.Scan(&countNumbers)
	if err != nil {
		fmt.Print("Ошибка ввода: ")
		fmt.Println(err)
	}

	var numbersArr []int
	fmt.Printf("Введите %v чисел по порядку:\n", countNumbers)
	var num int
	for i := 0; i < countNumbers; i++ {
		_, err = fmt.Scan(&num)
		if err != nil {
			fmt.Print("Ошибка ввода: ")
			fmt.Println(err)
		}
		numbersArr = append(numbersArr, num)
	}
	return numbersArr, countNumbers
}

func reverse(inputArr []int, countNumbers int) []int {
	var reverseArr []int
	for i := countNumbers - 1; i >= 0; i-- {
		reverseArr = append(reverseArr, inputArr[i])
	}
	return reverseArr
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Функция, реверсирующая массив.")
	fmt.Println()

	normalArr, countNumbers := inputNumbersNew()
	fmt.Printf("Массив в первичном порядке: %v\nМассив в обратном порядке: %v\n", normalArr, reverse(normalArr, countNumbers))
}
