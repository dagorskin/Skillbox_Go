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
					result[indexString][indexRune] = indexWord
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

	fmt.Println(finding(sentences, chars))

}
