package store

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}
