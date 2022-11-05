package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 7 (по желанию). Игра «Угадай число».")
	fmt.Println("------------------------------")
	fmt.Println("Добро пожаловать в игру \"Угадай число\"!")
	fmt.Println("Загадайте число от 1 до 10 (включительно), а программа попытается угадать.")
	fmt.Println("Если Ваше число больше используйте символ \">\", если меньше \"<\", а если программа угадала то \"=\".")

	// Инициализация переменных.
	var answer string
	countAttempts := 1
	number := 10

	// Угадываение числа.
	for true {
		fmt.Print("Это число ", number, "? Ваш ответ: ")
		fmt.Scan(&answer)
		if answer == "=" {
			fmt.Println("Программа угадала число с", countAttempts, "раза!")
			break
		} else if answer == ">" {
			countAttempts++
			if number > 0 && number <= 10 {
				number += (10 - number) / 2
				continue
			} else {
				fmt.Println("Вы вышли за границы!")
				continue
			}
		} else if answer == "<" {
			countAttempts++
			if number > 0 && number <= 10 {
				number /= 2
				continue
			} else {
				fmt.Println("Вы вышли за границы!")
				continue
			}
		}
	}
}
