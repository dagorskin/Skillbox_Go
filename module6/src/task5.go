package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 5 (по желанию). Задача на постепенное наполнение корзин разной ёмкости (if и continue и break).")
	fmt.Println("------------------------------")

	var firstBasketLimit, secondBasketLimit, thirdBasketLimit, firstBasket, secondBasket, thirdBasket, basket, fullBasket int

	// Ввод данных.
	fmt.Print("Введите ёмкость трёх корзин одной строкой через пробел: ")
	_, _ = fmt.Scan(&firstBasketLimit, &secondBasketLimit, &thirdBasketLimit)

	// Заполнение корзин.
	for {
		basket++
		// Первая корзина.
		if firstBasketLimit == basket {
			firstBasket = basket
			fullBasket++
			continue
		}
		// Вторая корзина.
		if secondBasketLimit == basket {
			secondBasket = basket
			fullBasket++
			continue
		}
		// Третья корзина.
		if thirdBasketLimit == basket {
			thirdBasket = basket
			fullBasket++
			continue
		}
		// Проверка на количество заполненных корзин.
		if fullBasket == 3 {
			break
		}
	}

	fmt.Print("Первая корзина заполнена на ", firstBasket, ", вторая на ", secondBasket, " и третья на ", thirdBasket, " яблок.")
}
