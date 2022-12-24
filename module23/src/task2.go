package main

import "fmt"

const arrayLength = 10

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Анонимные функции.")
	fmt.Println()

	var mainArray = [arrayLength]int{8, 5, 7, 1, 2, 4, 3, 9, 6, 10}

	// Первый вариант - сначала сортировка пузырьком, потом переворот массива.
	firstAnonymousFunction := func(intArray [arrayLength]int) [arrayLength]int {
		count := 0
	mainBreak:
		for {
			for i := 0; i < arrayLength-1; i++ {
				if intArray[i] > intArray[i+1] {
					intArray[i], intArray[i+1] = intArray[i+1], intArray[i]
					count++
				}
			}
			if count == 0 {
				break mainBreak
			}
			count = 0
		}

		// Переворот массива.
		for i, start, finish := 0, 0, arrayLength-1; i < arrayLength/2; i, start, finish = i+1, start+1, finish-1 {
			intArray[start], intArray[finish] = intArray[finish], intArray[start]
		}

		return intArray
	}

	// Второй вариант - сортировка ("вставками") с одновременным построением массива по убыванию.
	secondAnonymousFunction := func(intArray [arrayLength]int) [arrayLength]int {
		var digit, count int
		for index := 2; index < arrayLength; index++ {
			digit = intArray[index]
			count = index - 1
			for count >= 0 && intArray[count] < digit {
				intArray[count+1] = intArray[count]
				count -= 1
			}
			intArray[count+1] = digit
		}
		return intArray
	}

	resultFirst := firstAnonymousFunction(mainArray)
	resultSecond := secondAnonymousFunction(mainArray)

	fmt.Println("Начальный массив:\n", mainArray)
	fmt.Println("Массив после первого метода:\n", resultFirst)
	fmt.Println("Массив после второго метода:\n", resultSecond)
}
