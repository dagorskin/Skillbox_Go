package main

import "fmt"

// Инициализация размеров матриц.
const (
	mx3x5Row    = 3
	mx3x5Column = 5
	mx5x4Row    = 5
	mx5x4Column = 4
)

func matrixMultiplicationFirst(matrix3x5 [mx3x5Row][mx3x5Column]int, matrix5x4 [mx5x4Row][mx5x4Column]int) (result [3][4]int) {

	var i, j, l, k int
	var tempResult int

now:
	for i = 0; i < mx3x5Row; i++ {
		for j = 0; j < mx3x5Column; j++ {
			for l = 0; l < mx5x4Column; l++ {
				for k = 0; k < mx5x4Row; k++ {
					if j < mx3x5Row {
						tempResult += matrix3x5[j][k] * matrix5x4[k][l]
					} else {
						break now
					}
				}
				result[j][l] = tempResult
				tempResult = 0
			}
		}

	}
	return
}

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Умножение матриц.")
	fmt.Println()

	var matrix3x5 = [mx3x5Row][mx3x5Column]int{
		{5, 6, 8, 1, 7},
		{4, 1, 2, 3, 4},
		{9, 7, 6, 3, 1},
	}
	var matrix5x4 = [mx5x4Row][mx5x4Column]int{
		{3, 2, 4, 8},
		{6, 3, 9, 1},
		{1, 7, 2, 1},
		{4, 2, 3, 5},
		{5, 7, 4, 8},
	}

	fmt.Println(matrixMultiplicationFirst(matrix3x5, matrix5x4))
}
