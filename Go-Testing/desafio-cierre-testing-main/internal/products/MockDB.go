package products

import "fmt"

type MockDB struct {
	db  []Product
	err error
}

func (m *MockDB) GetAllBySeller(sellerID string) ([]Product, error) {
	if len(m.db) == 0 {
		return nil, fmt.Errorf("No products in db")
	}
	var prodList []Product
	for _, product := range m.db {
		if product.SellerID == sellerID {
			prodList = append(prodList, product)
		}
	}
	return prodList, nil
}
