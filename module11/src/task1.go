package main

import (
	"fmt"
	"strings"
)

func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Определение количества слов, начинающихся с большой буквы.")
	fmt.Println()

	fmt.Println("В тексте:")
	fmt.Println("<Go is an Open source programming Language that makes it Easy to build simple, " +
		"reliable, and efficient Software.>")

	var (
		countUppercase uint8
		wordAfter      string
		wordBefore     string
		spacePosition  int8
		textInput      = "Go is an Open source programming Language that makes it Easy to build simple, " +
			"reliable, and efficient Software"
	)

	for len(textInput) > 0 {
		textInput = strings.Trim(textInput, " ")
		spacePosition = int8(strings.Index(textInput, " "))

		if spacePosition == -1 {
			wordAfter = textInput
		} else {
			wordAfter = textInput[:spacePosition]
		}

		wordBefore = strings.Title(wordAfter)

		if wordBefore == wordAfter {
			countUppercase++
		}

		if spacePosition == -1 {
			break
		}

		textInput = textInput[spacePosition:]
	}
	fmt.Printf("Строка содержит %d слов с большой буквы.\n", countUppercase)
}
