package calculations

// Multiply Функция умножает число на цифру 2.
func Multiply(multiplyChan chan int) chan int {
	resultChan := make(chan int)
	number := <-multiplyChan
	result := number * 2
	go func() {
		resultChan <- result
	}()
	return resultChan
}
