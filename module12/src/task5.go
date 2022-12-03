package main

import (
	"fmt"
	"os"
	"strings"
)

func generatingBrackets(openBrackets, closedBrackets int, stringBrackets string, arrayBrackets *[]string) {
	if openBrackets == 0 {
		*arrayBrackets = append(*arrayBrackets, stringBrackets+strings.Repeat(")", closedBrackets))
		return
	}
	if closedBrackets > openBrackets {
		generatingBrackets(openBrackets, closedBrackets-1, stringBrackets+")", arrayBrackets)
	}
	generatingBrackets(openBrackets-1, closedBrackets, stringBrackets+"(", arrayBrackets)
}

func task5() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 5 (по желанию). Комбинации круглых скобок.")
	fmt.Println()

	file, err := os.OpenFile("task5.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Print("Файл не создан! Ошибка: ")
		fmt.Println(err)
	}
	defer file.Close()

	var quantityBrackets uint16
	fmt.Print("Введите количество пар скобок: ")
	_, _ = fmt.Scan(&quantityBrackets)

	closedBrackets, openBrackets := quantityBrackets, quantityBrackets
	var stringBrackets string
	arrayBrackets := make([]string, 0)

	generatingBrackets(int(closedBrackets), int(openBrackets), stringBrackets, &arrayBrackets)

	if _, err = file.WriteString(strings.Join(arrayBrackets, " ")); err != nil {
		fmt.Print("Ошибка при попытке записи: ")
		fmt.Println(err)
	}

	fmt.Println("Результат записан:\n", arrayBrackets)

}
