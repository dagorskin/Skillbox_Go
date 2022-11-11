package main

import "fmt"

func task3() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 3 (по желанию). Расчёт сдачи.")
	fmt.Println()

	// Ввод данных.
	const (
		fiveDollars   = 5
		tenDollars    = 10
		twentyDollars = 20
	)
	var lemonadeChange []int
	var bills, cashSum, change int

	for {
		fmt.Print("Введите купюру нового покупателя: ")
		_, _ = fmt.Scan(&bills)

		// Проверки на ввод.
		switch {
		case bills < fiveDollars:
			fmt.Println("Слишком маленькая сумма для покупки лимонада.")
		case bills-fiveDollars > cashSum:
			fmt.Println("В кассе нет столько денег чтобы дать сдачу.")
		case bills == fiveDollars:
			fmt.Println("Продажа прошла успешно, сдача ненужна.")
			cashSum += fiveDollars
			lemonadeChange = append(lemonadeChange, bills)
		case bills%fiveDollars == 0 || bills%tenDollars == 0 || bills%twentyDollars == 0:
			cashSum += fiveDollars
			lemonadeChange = append(lemonadeChange, bills)
			change = bills - fiveDollars
		default:
			fmt.Println("Такую сумму не разменять!")
		}

		// Проверки на "погашение" долга клиенту в виде сдачи.
	leaveSecondFor:
		for i := len(lemonadeChange) - 1; i >= 0; i-- {
			switch {
			case lemonadeChange[i] == twentyDollars && change >= twentyDollars:
				change -= twentyDollars
			case lemonadeChange[i] == tenDollars && change >= tenDollars:
				change -= tenDollars
			case lemonadeChange[i] == fiveDollars && change >= fiveDollars && change%fiveDollars == 0:
				change -= fiveDollars
			case change == 0:
				break leaveSecondFor
			}
		}

		// Вывод результата.
		fmt.Printf("В кассе %v долларов, а долг перед покупателем %v долларов;\n", cashSum, change)
		fmt.Println()

	}
}
