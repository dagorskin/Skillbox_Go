package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Генерация случайных чисел от -1000 до 1000.
func generatingRandom() int {

	t := time.NewTimer(time.Millisecond)
	rand.Seed(time.Now().UnixNano())
	randomNumber := -1000 + rand.Intn(1000-(-1000)+1)
	<-t.C
	return randomNumber
}

// Инициализация глобальных переменных.
var numberInput, result int
var varOne = generatingRandom()
var varTwo = generatingRandom()
var varThree = generatingRandom()

func summationFirst(number, variable int) int {
	return number + variable
}

func summationSecond(number, variable int) int {
	return number + variable
}

func summationThird(number, variable int) int {
	return number + variable
}

func task4() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 4. Область видимости переменных.")
	fmt.Println()

	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&numberInput)
	if err != nil {
		fmt.Print("Ошибка ввода :")
		fmt.Println(err)
	}

	result = summationThird(summationSecond(summationThird(numberInput, varOne), varTwo), varThree)

	fmt.Println("Использованные переменные:")
	fmt.Println(varOne, varTwo, varThree)
	fmt.Printf("Результат: %d;\n", result)

}
