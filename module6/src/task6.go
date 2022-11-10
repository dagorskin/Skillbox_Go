package main

import "fmt"

func main() {
	fmt.Println("------------------------------")
	fmt.Println("Задание 6 (по желанию). Движение лифта.")
	fmt.Println("------------------------------")

	// Инициализация заведомо известных данных.
	const groundFloor, lastFloor, weightLimit = 1, 24, 2
	var floors []int
	elevatorPosition, numberPeopleInElevator, quantityPeopleGroundFloor, quantityPeopleTheFloors := 1, 0, 0, 0

	// Основной цикл работы программы.
	for {
		elevatorPosition = groundFloor
		fmt.Println("Лифт на", elevatorPosition, "этаже.")

		// Ввод данных о людях нуждающихся в услугах лифта.
		if numberPeopleInElevator == 0 && quantityPeopleGroundFloor == 0 && quantityPeopleTheFloors == 0 {
			fmt.Print("Сколько человек хотят воспользоваться лифтом на первом этаже? ")
			_, _ = fmt.Scan(&quantityPeopleGroundFloor)
			fmt.Print("Сколько человек хотят воспользоваться лифтом других этажах? ")
			_, _ = fmt.Scan(&quantityPeopleTheFloors)

			//  Добавление людей на этажах в срез.
			if quantityPeopleTheFloors != 0 {
				fmt.Print("Перечислите эти этажи: ")
				for i := quantityPeopleTheFloors; i > 0; i-- {
					var tempPeople int
					_, _ = fmt.Scan(&tempPeople)
					floors = append(floors, tempPeople)
				}
			}
		}

		// Посадка пассажиров в лифт.
		numberPeopleInElevator = 0
		if quantityPeopleGroundFloor >= weightLimit {
			numberPeopleInElevator += 2
			quantityPeopleGroundFloor -= 2
		} else if quantityPeopleGroundFloor == 0 && quantityPeopleTheFloors == 0 {
			fmt.Print("В лифте осталось ", numberPeopleInElevator, ", на первом этаже ", quantityPeopleGroundFloor, ".\n")
			fmt.Println("Лифт развёз всех людей!")
			break
		} else {
			numberPeopleInElevator = quantityPeopleGroundFloor
			quantityPeopleGroundFloor = 0
		}

		// Движение вверх.
		if quantityPeopleGroundFloor > 0 || quantityPeopleTheFloors > 0 {
			fmt.Print("В лифт загрузилось ", numberPeopleInElevator, " человек, осталось на первом этаже ", quantityPeopleGroundFloor, ".\n")
			fmt.Println("Лифт движется наверх.")
			elevatorPosition = lastFloor
			fmt.Println("Лифт на", elevatorPosition, "этаже. Все кто были в лифте вышли.")
			numberPeopleInElevator = 0
			fmt.Print("В лифте осталось ", numberPeopleInElevator, ", на первом этаже ", quantityPeopleGroundFloor, ".\n")
		}

		// Движение вниз.
		fmt.Println("Лифт пустой и движется вниз.")
		for quantityPeopleTheFloors > 0 {
			elevatorPosition = floors[quantityPeopleTheFloors-1]
			quantityPeopleTheFloors--
			numberPeopleInElevator++
			floors = append(floors[:quantityPeopleTheFloors], floors[quantityPeopleTheFloors+1:]...)
			fmt.Println("Лифт остановился на", elevatorPosition, "этаже, забрал человека и поехал дальше.")
			fmt.Println("В лифте теперь:", numberPeopleInElevator, "человек.")

			if numberPeopleInElevator == 2 {
				fmt.Println("Лифт полон и движется вниз не подбирая никого.")
				fmt.Println("---------------------------------------------")
				break
			}
		}
	}
}
