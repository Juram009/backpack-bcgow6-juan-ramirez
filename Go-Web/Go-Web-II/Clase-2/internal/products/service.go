package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name, color string, price, stock int, code string, published bool, creationDate string) (Product, error)
	Update(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error)
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

func (s *service) Store(name, color string, price, stock int, code string, published bool, creationDate string) (Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, name, color, price, stock, code, published, creationDate)
	if err != nil {
		return Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, name, color string, price, stock int, code string, published bool, creationDate string) (Product, error) {

	return s.repository.Update(id, name, color, price, stock, code, published, creationDate)
}
