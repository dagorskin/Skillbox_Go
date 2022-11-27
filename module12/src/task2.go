package main

import (
	"fmt"
	"io"
	"os"
)

func task2() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 2. Интерфейс io.Reader.")
	fmt.Println()

	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла!", err)
		return
	}
	defer file.Close()

	fileSize, _ := file.Stat()
	buffer := make([]byte, fileSize.Size())
	if _, err = io.ReadFull(file, buffer); err != nil {
		fmt.Println("Не получилось прочитать последовательность файлов!", err)
		return
	}

	fmt.Printf("Вывод информации с файла: \n%s\n", buffer)
}
