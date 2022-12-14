package main

import "fmt"

// Инициализация размеров матриц.
const (
	mx3x5Row    = 3
	mx3x5Column = 5
	mx5x4Row    = 5
	mx5x4Column = 4
)

func matrixMultiplicationFirst(matrix3x5 [mx3x5Row][mx3x5Column]int, matrix5x4 [mx5x4Row][mx5x4Column]int) (result [mx3x5Row][mx5x4Column]int) {

	var j, k, l int
	var tempResult int

now:
	for j = 0; j < mx3x5Column; j++ {
		for k = 0; k < mx5x4Column; k++ {
			for l = 0; l < mx5x4Row; l++ {
				if j < mx3x5Row {
					tempResult += matrix3x5[j][l] * matrix5x4[l][k]
				} else {
					break now
				}
			}
			result[j][k] = tempResult
			tempResult = 0
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

	var result = matrixMultiplicationFirst(matrix3x5, matrix5x4)

	fmt.Println("Результат умножения матриц: ")
	for row := 0; row < len(result); row++ {
		fmt.Println(result[row])
	}
}
