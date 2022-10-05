package products

type Service interface {
	GetAll() ([]Product, error)
	GetOne(id int) (Product, error)
	Store(name, color string, price, stock int, code string, published bool, creationDate string) (Product, error)
	Update(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error)
	Delete(id int) error
	UpdateNamePrice(id int, name string, price int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetOne(id int) (Product, error) {
	ps, err := s.repository.GetOne(id)
	if err != nil {
		return Product{}, err
	}
	return ps, nil
}

func (s *service) Store(name, color string, price, stock int, code string, published bool, creationDate string) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	producto, err := s.repository.Store(lastID+1, name, color, price, stock, code, published, creationDate)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error) {
	return s.repository.Update(id, name, color, price, stock, code, published, creationDate)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateNamePrice(id int, name string, price int) (Product, error) {
	return s.repository.UpdateNamePrice(id, name, price)
}
