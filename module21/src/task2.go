package main

import "fmt"

func anonymous(numberOne float64, numberTwo float64, A func (float64, float64) float64) {
	defer fmt.Println(A(numberOne, numberTwo))
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Анонимные функции.")
	fmt.Println()

	var numberOne, numberTwo float64
	fmt.Print("Введите два значения через пробел: ")
	_, _ = fmt.Scan(&numberOne, & numberTwo)

	anonymous(numberOne, numberTwo, func(float64,float64) float64 {return numberOne + numberTwo})
	anonymous(numberOne, numberTwo, func(float64,float64) float64 {return numberOne - numberTwo})
	anonymous(numberOne, numberTwo, func(float64,float64) float64 {return numberOne * numberTwo})
	anonymous(numberOne, numberTwo, func(float64,float64) float64 {return numberOne / numberTwo})
}
