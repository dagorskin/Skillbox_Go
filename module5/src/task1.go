package main

import "fmt"

// Объявление глобальных переменных.
var pointX, pointY int

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Определение координатной плоскости точки.")
	fmt.Println("------------------------------")
	// Ввод данных.
	fmt.Print("Введите точку X: ")
	fmt.Scan(&pointX)
	fmt.Print("Введите точку Y: ")
	fmt.Scan(&pointY)

	// Нахождение точки и вывод результата.
	if pointX > 0 && pointY > 0 {
		fmt.Println("Точка находится в I четверти координатной плоскости.")
	} else if pointX < 0 && pointY > 0 {
		fmt.Println("Точка находится в II четверти координатной плоскости.")
	} else if pointX < 0 && pointY < 0 {
		fmt.Println("Точка находится в III четверти координатной плоскости.")
	} else if pointX > 0 && pointY < 0 {
		fmt.Println("Точка находится в IV четверти координатной плоскости.")
	} else if pointX == 0 && pointY == 0 {
		fmt.Println("Точка находится на пересечении координатной плоскости в точке O.")
	} else if pointX == 0 && pointY != 0 {
		fmt.Println("Точка находится на координатной оси Y")
	} else if pointY == 0 && pointX != 0 {
		fmt.Println("Точка находится на координатной оси X")
	} else {
		fmt.Println("Нет возможности определить нахождение точки.")
	}
}
