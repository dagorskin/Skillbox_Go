package main

import "fmt"

func main() {
	var answer string

	for {
		fmt.Println("------------------------------")
		fmt.Print("Введите номер задания для его запуска (Q - выход): ")
		fmt.Scan(&answer)

		switch answer {
		case "1":
			task1()
		case "2":
			task2()
		case "3":
			task3()
		case "4":
			task4()
		case "Q":
			return
		case "q":
			return
		}
	}
}
