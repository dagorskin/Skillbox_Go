package main

import (
	"fmt"
	"task1/pkg/calculations"
	"task1/pkg/input"
	"task1/pkg/storage"
)

// App Структура программы.
type App struct {
	repository storage.Storage
}

// NewApp Конструктор структуры программы. Возвращает новый объект.
func NewApp(repository storage.Storage) *App {
	return &App{repository: repository}
}

// Run Функция запуска программы.
func (a *App) Run() {
	squareChan := make(chan int)

	for {
		if number, ok := input.Input(); ok {
			go func() {
				squareChan <- number
			}()
			squareRes, sr := calculations.Square(squareChan)
			mp := calculations.Multiply(sr)
			stringData := storage.DataString{InputNum: number, SquareRes: squareRes,
				MultiplyRes: <-mp}
			a.Store(stringData)
		} else {
			a.Print()
			break
		}
	}
}

// Print Функция вывода результата.
func (a *App) Print() {
	data := a.repository.Get()
	fmt.Println("\nВвод окончен. Вывод результата:")
	for _, value := range data {
		fmt.Printf("%v. Число: %v квадрат: %v произведение: %v\n",
			value.Position, value.InputNum, value.SquareRes, value.MultiplyRes)
	}
}

// Store Функция добавления числа в строку карты методом Put.
func (a *App) Store(number storage.DataString) {
	a.repository.Put(number)
}

// Входная точка.
func main() {
	repository := storage.NewMemStorage()

	var app = NewApp(repository)

	app.Run()

	fmt.Println("...вывод закончен.")
}
