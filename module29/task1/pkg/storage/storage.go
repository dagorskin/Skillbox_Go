package storage

// Storage Интерфейс хранилища.
type Storage interface {
	Put(DataString)
	Get() map[int]DataString
}
