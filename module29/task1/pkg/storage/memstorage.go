package storage

var position = 1 // Позиция строки в карте.

// DataString Структура строки в карте.
type DataString struct {
	Position    int
	InputNum    int
	MultiplyRes int
	SquareRes   int
}

// MemStorage Карта, расположенная в хранилище.
type MemStorage struct {
	Data map[int]DataString
}

// NewMemStorage Конструктор карты. Возвращает готовую карту.
func NewMemStorage() *MemStorage {
	return &MemStorage{
		Data: make(map[int]DataString),
	}
}

// Put Функция записи строки в карту.
func (ms *MemStorage) Put(data DataString) {
	data.Position = position
	ms.Data[position] = data
	position++
}

// Get Функция получения карты из хранилища.
func (ms *MemStorage) Get() map[int]DataString {
	return ms.Data
}
