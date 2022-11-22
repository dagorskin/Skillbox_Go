package main

import (
	"fmt"
	"strconv"
	"strings"
)

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Вывод чисел, содержащихся в строке.")
	fmt.Println()

	// Инициализация.
	var (
		textInput  = "a10 10 20b 20 30c30 30 dd"
		textOutput string
		tempString string
	)

	// Основной цикл.
	for i := 0; i < len(textInput); i++ {
		if textInput[i] != ' ' {
			tempString += string(textInput[i])
		} else {
			_, err := strconv.ParseInt(tempString, 10, 8)
			if err != nil {
				textInput = textInput[i:]
				tempString = ""
				i = 0
				continue
			} else {
				textOutput += tempString + " "
				tempString = ""
			}
		}
	}

	fmt.Println("В строке содержатся числа в десятичном формате:")
	fmt.Println(strings.Trim(textOutput, " "))
}
