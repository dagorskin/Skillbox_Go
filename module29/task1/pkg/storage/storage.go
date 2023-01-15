package storage

// Storage Интерфейс хранилища.
type Storage interface {
	Put(DataString)
	Get() []DataString
}
