package products

import "github.com/Juram009/backpack-bcgow6-juan-ramirez/Go-Testing/Go-Testing-II/clase-2/clase/internal/domain"

type MockStorage struct {
	dataMock   []domain.Product
	errOnWrite error
	errOnRead  error
}

func (m *MockStorage) Read(data interface{}) (err error) {
	if m.errOnRead != nil {
		return m.errOnRead
	}

	castedData := data.(*[]domain.Product)
	*castedData = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}

	castedData := data.(*domain.Product)
	m.dataMock = append(m.dataMock, *castedData)
	return nil
}
