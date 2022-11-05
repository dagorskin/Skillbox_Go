package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 5. Решение квадратного уравнения.")
	fmt.Println("------------------------------")

	// Инициализация переменных.
	var coefficientA, coefficientB, coefficientC, discriminant, root1, root2 float64

	// Ввод данных.
	fmt.Println("Вас приветствует программа для вычисления квадратного уравнения под названием <Решала>!")
	fmt.Print("Введите коэффициент А: ")
	fmt.Scan(&coefficientA)
	fmt.Print("Введите коэффициент B: ")
	fmt.Scan(&coefficientB)
	fmt.Print("Введите коэффициент C: ")
	fmt.Scan(&coefficientC)

	// Вычисление дискриминанта.
	discriminant = math.Pow(coefficientB, 2) - 4*coefficientA*coefficientC

	// Нахождение корней и вывод их значений.
	if discriminant > 0 {
		root1 = (-coefficientB - math.Sqrt(discriminant)) / (2 * coefficientA)
		root2 = (-coefficientB + math.Sqrt(discriminant)) / (2 * coefficientA)
		fmt.Print("Так как дискриминант больше нуля, то корня два: ", root1, " и ", root2, ".")
	} else if discriminant == 0 {
		root1 = (-coefficientB) / 2 * coefficientA
		fmt.Print("Так как дискриминант больше равен нулю, то корень один: ", root1, ".")
	} else if discriminant < 0 {
		fmt.Println("Дискриминант меньше нуля - корней нет!")
	}
}
