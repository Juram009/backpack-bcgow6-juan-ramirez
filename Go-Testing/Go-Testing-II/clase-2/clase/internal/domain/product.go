package domain

// Product structure used to store product information.
type Product struct {
	ID    int     `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Type  string  `json:"type,omitempty"`
	Count int     `json:"count,omitempty"`
	Price float64 `json:"price,omitempty"`
}
