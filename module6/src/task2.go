package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Сумма двух чисел по единице.")
	fmt.Println("------------------------------")

	var firstNum, secondNum int

	// Ввод данных.
	fmt.Print("Введите первое число: ")
	_, _ = fmt.Scan(&firstNum)
	fmt.Print("Введите второе число: ")
	_, _ = fmt.Scan(&secondNum)

	// Сам цикл.
	sumNum := firstNum + secondNum
	fmt.Print("Сумма чисел равна: ", sumNum, ".\n")
	for firstNum != sumNum {
		firstNum++
		fmt.Print("Первое число: ", firstNum, "; \n")
	}

	// Вывод сообщения о завершении циклы.
	fmt.Print("Цикл закончен: первое число выросло до = ", firstNum, " и теперь ровно сумме введенных чисел.")
}
