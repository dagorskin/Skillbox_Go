package main

import (
	"fmt"
	"strings"
)

func finding(sentences [4]string, chars [5]rune) (result [4][5]int) {
	for indexString, letter := range sentences {
		for indexWord, word := range letter {
			for indexRune, char := range chars {
				if strings.ToLower(string(char)) == strings.ToLower(string(word)) {
					result[indexString][indexRune] = indexWord + 1
				}
			}
		}
	}
	return
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Поиск символов в нескольких строках.")
	fmt.Println()

	sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [5]rune{'H', 'E', 'L', 'П', 'М'}

	result := finding(sentences, chars)
	fmt.Println("Результат поиска букв в строках:\n", result)

	fmt.Println("\nРезультат первого элемента:")
	for i := 0; i != 5; i++ {
		if result[0][i] == 0 {
			continue
		} else {
			fmt.Printf("%q позиция %d\n", chars[i], result[0][i]-1)
		}
	}

	fmt.Println("\nРезультат второго элемента:")
	for i := 0; i != 5; i++ {
		if result[1][i] == 0 {
			continue
		} else {
			fmt.Printf("%q позиция %d\n", chars[i], result[1][i]-1)
		}
	}

	fmt.Println("\nРезультат третьего элемента:")
	for i := 0; i != 5; i++ {
		if result[2][i] == 0 {
			continue
		} else {
			fmt.Printf("%q позиция %d\n", chars[i], result[2][i]-1)
		}
	}

	fmt.Println("\nРезультат четвертого элемента:")
	for i := 0; i != 5; i++ {
		if result[3][i] == 0 {
			continue
		} else {
			fmt.Printf("%q позиция %d\n", chars[i], result[3][i]-1)
		}
	}

}
