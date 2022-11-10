package main

import "fmt"

func task3() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 3. Вывод ёлочки.")
	fmt.Println()

	// Инициализация данных разных типов.
	var heightTree, weightTree int

	// Ввод данных.
	fmt.Print("Введите высоту ёлочки: ")
	_, _ = fmt.Scan(&heightTree)
	weightTree = heightTree*2 - 1
	before := heightTree - 2
	after := weightTree - (heightTree - 1)

	// Рисование "ёлочки"
	for vertical := heightTree; vertical != 0; vertical-- {
		for horizontal := 0; horizontal < weightTree; horizontal++ {
			if horizontal <= before {
				fmt.Print(" ")
			}
			if horizontal > before && horizontal < after {
				fmt.Print("*")
			}
			if horizontal >= after {
				fmt.Print(" ")
			}
		}
		before--
		after++
		fmt.Println("")
	}
}
