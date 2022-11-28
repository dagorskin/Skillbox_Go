package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func task4() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 4. Пакет ioutil.")
	fmt.Println()

	var (
		buffer    bytes.Buffer
		fileName  = "strings.txt"
		dataWrite string
	)

	answer := ""
	fmt.Print("Введите строку для записи в файл: ")
	_, _ = fmt.Scan(&answer)

	timeNow := time.Now().Format("2006-01-02 15:04:05")
	dataWrite = timeNow + " " + answer + "\n"

	_, err := buffer.WriteString(dataWrite)
	if err != nil {
		panic(err)
	}

	if err = ioutil.WriteFile(fileName, buffer.Bytes(), 0666); err != nil {
		panic(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Вы ввели: \n%s\n", resultBytes)

}
