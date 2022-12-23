package main

import "fmt"

func finding(sentences [4]string, chars [5]rune) (result [5][5]int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			for index, char := range sentences[j] {
				if chars[i] == char {
					result[i][j] = index
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
