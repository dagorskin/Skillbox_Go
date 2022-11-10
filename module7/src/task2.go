package main

import "fmt"

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Шахматная доска.")
	fmt.Println()

	// Инициализация данных разных типов.
	var widthChessboard, heightChessboard, square int
	// Ввод данных.
	fmt.Print("Введите ширину шахматной доски: ")
	_, _ = fmt.Scan(&widthChessboard)
	fmt.Print("Введите высоту шахматной доски: ")
	_, _ = fmt.Scan(&heightChessboard)

	// Рисование "шахматной доски".
	for vertical := heightChessboard; vertical != 0; vertical-- {
		for horizontal := widthChessboard; horizontal != 0; horizontal-- {
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
