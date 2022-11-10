package main

import "fmt"

func task3() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 3 (по желанию). Расчёт сдачи.")
	fmt.Println()

	// Ввод данных.
	const (
		fiveRuble   = 5
		tenRuble    = 10
		twentyRuble = 20
	)
	var lemonadeChange []int
	var bills, cashSum, change int

	for {
		fmt.Print("Введите купюру нового покупателя: ")
		_, _ = fmt.Scan(&bills)

		// Проверки на ввод.
		switch {
		case bills < fiveRuble:
			fmt.Println("Слишком маленькая сумма для покупки лимонада.")
		case bills-fiveRuble > cashSum:
			fmt.Println("В кассе нет столько денег чтобы дать сдачу.")
		case bills == fiveRuble:
			fmt.Println("Продажа прошла успешно, сдача ненужна.")
			cashSum += fiveRuble
			lemonadeChange = append(lemonadeChange, bills)
		case bills%fiveRuble == 0 || bills%tenRuble == 0 || bills%twentyRuble == 0:
			cashSum += fiveRuble
			lemonadeChange = append(lemonadeChange, bills)
			change = bills - fiveRuble
		default:
			fmt.Println("Такую сумму не разменять!")
		}

		// Проверки на "погашение" долга клиенту в виде сдачи.
	leaveSecondFor:
		for i := len(lemonadeChange) - 1; i >= 0; i-- {
			switch {
			case lemonadeChange[i] == twentyRuble && change >= twentyRuble:
				change -= twentyRuble
			case lemonadeChange[i] == tenRuble && change >= tenRuble:
				change -= tenRuble
			case lemonadeChange[i] == fiveRuble && change >= fiveRuble && change%fiveRuble == 0:
				change -= fiveRuble
			case change == 0:
				break leaveSecondFor
			}
		}

		// Вывод результата.
		fmt.Printf("В кассе %v рублей, а долг перед покупателем %v рублей;\n", cashSum, change)
		fmt.Println()

	}
}
