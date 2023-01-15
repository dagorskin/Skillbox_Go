package calculations

// Square Функция высчитывания квадрата числа.
func Square(squareChan chan int) (int, chan int) {
	multiplyChan := make(chan int)
	number := <-squareChan
	result := number * number
	go func() {
		multiplyChan <- result
	}()
	return result, multiplyChan
}
