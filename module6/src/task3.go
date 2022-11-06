package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 3. Расчёт суммы скидки.")
	fmt.Println("------------------------------")

	var priceInput, discountInput int

	// Ввод данных.
	fmt.Print("Введите стоимость товара: ")
	fmt.Scan(&priceInput)
	fmt.Print("Введите размер скидки в процентах: ")
	fmt.Scan(&discountInput)

	discount := discountInput * priceInput / 100

	// Ветвление с выводом результата.
	if discount < 2000 || discountInput < 30 {
		fmt.Println("Ваша скидка равна:", discount, "рублей.")
	} else {
		fmt.Println("Ваша скидка максимальная и она равна: 2000 рублей.") // Верхняя граница суммы 2000 рублей.
	}
}
