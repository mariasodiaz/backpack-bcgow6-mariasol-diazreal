package products

type Service interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(name string, color string, price int, stock int, code string, published bool, date string) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Store(name string, color string, price int, stock int, code string, published bool, date string) (Product, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return Product{}, err
	}
	id++
	product, err := s.repository.Store(id, name, color, price, stock, code, published, date)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) GetById(id int) (Product, error) {
	product, err := s.repository.GetById(id)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
