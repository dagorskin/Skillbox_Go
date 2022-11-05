package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 4. Сумма без сдачи.")
	fmt.Println("------------------------------")

	// Инициализация переменных.
	var costGoods, firstCoin, secondCoin, thirdCoin int

	// Ввод данных.
	fmt.Println("Здравствуйте! Вас приветствует терминал банка <СкиллБанк>")
	fmt.Print("Введите сумму для снятия: ")
	fmt.Scan(&costGoods)
	fmt.Print("Введите номиналы трёх монет через запятую: ")
	fmt.Scan(&firstCoin, &secondCoin, &thirdCoin)

	sumCoin := firstCoin + secondCoin + thirdCoin

	// Условия выполнения программы с тем или иным выводом.
	if costGoods == firstCoin || costGoods == secondCoin || secondCoin == thirdCoin || costGoods == sumCoin || costGoods%firstCoin == 0 || costGoods%secondCoin == 0 || costGoods%thirdCoin == 0 {
		fmt.Println("Вы можете получит эту сумму и оплатить товар!")
	} else {
		fmt.Println("Вы не можете получить нужную сумму!")
	}
}
