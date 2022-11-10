package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 4. Предыдущее занятие на if.")
	fmt.Println("------------------------------")

	var firstVar, secondVar, thirdVar int

	fmt.Print("Первое число = ", firstVar, ", второе = ", secondVar, ", третье = ", thirdVar, ".\n")

	// Первый вариант.
	/*for {
		if firstVar != 10 {
			firstVar++
		}
		if secondVar != 100 {
			secondVar++
		}
		if thirdVar != 1000 {
			thirdVar++
		}
		if firstVar == 10 && secondVar == 100 && thirdVar == 1000 {
			break
		}
	}*/

	// Второй вариант.
	/*for thirdVar != 1000 {
		if firstVar != 10 {
			firstVar++
		} else if secondVar != 100 {
			secondVar++
		} else if thirdVar != 1000 {
			thirdVar++
		}
	}*/

	// Третий вариант.
	for thirdVar != 1000 {
		thirdVar++
		for secondVar != 100 {
			secondVar++
			for firstVar != 10 {
				firstVar++
			}
		}
	}

	fmt.Print("Инкрементирование закончена: первое число = ", firstVar, ", второе = ", secondVar, " и третье число = ", thirdVar, ".")

}
