package main

import "fmt"

const rows, columns = 3, 3

// Сначала расширил матрицу.
func matrixExpansion(matrix3x3 [rows][columns]int) (matrix5x5 [3][5]int) {
	for row := 0; row < len(matrix3x3); row++ {
		for column := 0; column < len(matrix3x3[row]); column++ {
			matrix5x5[row][column] = matrix3x3[row][column]
			if column == 2 {
				matrix5x5[row][column+1] = matrix3x3[row][0]
				matrix5x5[row][column+2] = matrix3x3[row][1]
			}
		}
	}
	return
}

// Произведения левых элементов сложил, потом вычитал результат произведения правых элементов.
func ruleSarrus(matrix3x5 [3][5]int) (totalDeterminant int) {
	determinant := 1
	countColumn := 0
	for i := 0; i < len(matrix3x5); i++ {
		for row := 0; row < len(matrix3x5); row++ {
			for column := 0; column < len(matrix3x5[row]); column++ {
				if row == column-countColumn {
					determinant = determinant * matrix3x5[row][column]
				}
			}
		}
		countColumn++
		totalDeterminant += determinant
		determinant = 1
	}

	for i := 0; i < len(matrix3x5); i++ {
		countColumn = 4 - i
		for row := 0; row < len(matrix3x5); row++ {
			for column := 0; column < len(matrix3x5[row]); column++ {
				if row == column-countColumn {
					determinant = determinant * matrix3x5[row][column]
				}
			}
			countColumn -= 2
		}
		totalDeterminant -= determinant
		determinant = 1
	}
	return
}

// Для решения применил Правило Саррюса.
func task1() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 1. Подсчёт определителя.")
	fmt.Println()

	var matrix3x3 = [rows][columns]int{
		{1, 2, 4},
		{6, 8, 7},
		{5, 9, 3},
	}

	fmt.Println("Определитель матрицы: ", ruleSarrus(matrixExpansion(matrix3x3)))
}
