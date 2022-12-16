package main

import (
	"fmt"
	"math/rand"
	"time"
)

const length = 10

// Создание массива из оригинальных случайных цифр.
func generationArrUnsorted() (array [length] int) {
	rand.Seed(time.Now().UnixNano())
	var tempNumber, count int
	for index := 0; index < length; index++ {
		tempNumber = rand.Intn(length+(length*0.5))
		for _, number := range array {
			if number == tempNumber {
				count++
			}
		}
		if count == 0 {
			array[index] = tempNumber
		} else {
			index--
		}
		count = 0
	}
	return
}

// Поиск числа и вывод результата.
func numberSearch(array[length] int, desiredNumber int) (result int){
	for index, number := range array {
		if number == desiredNumber {
			result = length-index-1
			break
		}
	}
	return
}

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Подсчёт чисел в массиве.")
	fmt.Println()

	var desiredNumber, result int
	fmt.Print("Введите число: ")
	_, _ = fmt.Scan(&desiredNumber)

	array := generationArrUnsorted()
	result = numberSearch(array, desiredNumber)
	fmt.Printf("Используемый массив:\n%v\n", array)
	fmt.Printf("Результат: %v\n", result)
}