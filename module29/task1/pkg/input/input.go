package input

import (
	"fmt"
	"strconv"
)

// Input Вернет вводимое число и нужно ли продолжать ввод.
func Input() (int, bool) {
	var stringTemp string
	var number int

	fmt.Print("Введите число (стоп - завершить ввод): ")
	_, err := fmt.Scan(&stringTemp)

	if stringTemp == "стоп" || stringTemp == "Стоп" || stringTemp == "СТОП" {
		return 0, false
	}

	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		fmt.Print("Повторите ввод. ")
		Input()
	} else {
		number, err = strconv.Atoi(stringTemp)
		if err != nil {
			fmt.Println("Ошибка ввода: ", err)
			Input()
		}
	}
	return number, true
}
