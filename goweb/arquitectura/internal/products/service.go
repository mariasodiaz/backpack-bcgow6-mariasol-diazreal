package products

type Service interface {
	GetAll() ([]Product, error)
	Store(name string, color string, price int, stock int, code string, published bool, date string) (Product, error)
	Update(id int, name string, color string, price int, stock int, code string, published bool, date string) (Product, error)
	Delete(id int) error
	UpdateMany(id int, name string, price int) (Product, error)
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

func (s *service) Update(id int, name string, color string, price int, stock int, code string, published bool, date string) (Product, error) {
	product, err := s.repository.Update(id, name, color, price, stock, code, published, date)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) UpdateMany(id int, name string, price int) (Product, error) {
	product, err := s.repository.UpdateMany(id, name, price)
	if err != nil {
		return Product{}, err
	}
	return product, nil
}
