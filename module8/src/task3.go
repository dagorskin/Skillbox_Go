package main

import "fmt"

func task3() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 3 (по желанию). Расчёт сдачи.")
	fmt.Println()

	// Ввод данных.
	var lemonadeChange []int
	var bills, cashSum, change int

	for {
		fmt.Print("Введите купюру покупателя: ")
		_, _ = fmt.Scan(&bills)

		switch {
		case bills < 5:
			fmt.Println("Слишком маленькая сумма для покупки лимонада.")
		case bills-5 > cashSum:
			fmt.Println("В кассе нет столько денег чтобы дать сдачу.")
		case bills == 5:
			fmt.Println("Продажа прошла успешно, сдача ненужна.")
			cashSum += 5
			lemonadeChange = append(lemonadeChange, bills)
		case bills%5 == 0 || bills%10 == 0 || bills%20 == 0:
			cashSum += 5
			lemonadeChange = append(lemonadeChange, bills)
			change = bills - 5
		}

	leaveSecondFor:
		for {
			switch {
			case change >= 20 && change%20 == 0:
				change -= 20
			case change >= 10 && change%10 == 0:
				change -= 10
			case change >= 5 && change%5 == 0:
				change -= 5
			case change == 0:
				break leaveSecondFor
			}
		}

		fmt.Println(cashSum, change)

	}
}
