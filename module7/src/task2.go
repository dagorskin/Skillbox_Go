package main

import "fmt"

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Шахматная доска.")
	fmt.Println("------------------------------")

	// Инициализация данных разных типов.
	var widthСhessboard, heightСhessboard, square int
	// Ввод данных.
	fmt.Print("Введите ширину шахматной доски: ")
	fmt.Scan(&widthСhessboard)
	fmt.Print("Введите высоту шахматной доски: ")
	fmt.Scan(&heightСhessboard)

	// Рисование шахматной доски.
	for vertical := heightСhessboard; vertical != 0; vertical-- {
		for horizontal := widthСhessboard; horizontal != 0; horizontal-- {
			square++
			if square%2 == 0 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
