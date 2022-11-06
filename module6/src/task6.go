package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 6 (по желанию). Движение лифта.")
	fmt.Println("------------------------------")

	const groundFloor, lastFloor, weightLimit = 1, 24, 2
	var people []int
	elevatorPosition, numberPeopleInElevator, quantityPeopleGroundFloor, quantityPeopleTheFloors := 1, 0, 0, 0

	for {
		fmt.Println("Лифт на", elevatorPosition, "этаже.")

		if elevatorPosition == 1 {
			fmt.Print("Сколько человек хотят воспользоваться лифтом на первом этаже? ")
			fmt.Scan(&quantityPeopleGroundFloor)

			// Посадка пассажиров в лифт.
			if quantityPeopleGroundFloor <= 2 {
				numberPeopleInElevator = quantityPeopleGroundFloor
			} else {
				numberPeopleInElevator += 2
				quantityPeopleGroundFloor -= 2
			}

			fmt.Print("Сколько человек хотят воспользоваться лифтом других этажах? ")
			fmt.Scan(&quantityPeopleTheFloors)

			//  Добавление людей на этажах в срез.
			if quantityPeopleTheFloors != 0 {
				fmt.Print("Перечислите эти этажи: ")
				for i := quantityPeopleTheFloors; i > 0; i-- {
					var tempPeople int
					fmt.Scan(&tempPeople)
					people = append(people, tempPeople)
				}
			}
			// Движение вверх.
			if quantityPeopleGroundFloor > 0 {
				fmt.Println("Лифт движется наверх развозить людей.")
				elevatorPosition = 24
				fmt.Println("Лифт на", elevatorPosition, "этаже.")
				numberPeopleInElevator = 0
				fmt.Print("В лифте осталось ", numberPeopleInElevator, ", а на первом этаже ", quantityPeopleGroundFloor, ".\n")
			}
		}
	}
}
