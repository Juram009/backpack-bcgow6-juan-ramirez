package products

type MockStorage struct {
	dataMock      []Product
	errOnWrite    error
	errOnRead     error
	ReadWasCalled bool
}

func (m *MockStorage) Read(data interface{}) (err error) {
	if m.errOnRead != nil {
		return m.errOnRead
	}
	castedData := data.(*[]Product)
	*castedData = m.dataMock
	m.ReadWasCalled = true
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}

	castedData := data.([]Product)
	m.dataMock = castedData
	return nil
}
