package products

import (
	"errors"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/internal/domain"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/pkg/store"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, name string, color string, price int, stock int, code string, published bool, date string) (domain.Product, error)
	LastId() (int, error)
	Update(id int, name string, color string, price int, stock int, code string, published bool, date string) (domain.Product, error)
	Delete(id int) error
	UpdateMany(id int, name string, price int) (domain.Product, error)
}

type repository struct {
	db store.Store
}

var errorIdNotFound = errors.New("id not found")

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	var products = []domain.Product{}
	r.db.Read(&products)
	return products, nil
}

func (r *repository) LastId() (int, error) {
	var products = []domain.Product{}
	if err := r.db.Read(&products); err != nil {
		return 0, err
	}
	if len(products) == 0 {
		return 0, nil
	}
	return products[len(products)-1].Id, nil
}

func (r *repository) Store(id int, name string, color string, price int, stock int, code string, published bool, date string) (domain.Product, error) {
	product := domain.Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}
	var products = []domain.Product{}
	r.db.Read(&products)
	products = append(products, product)
	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) Update(id int, name string, color string, price int, stock int, code string, published bool, date string) (domain.Product, error) {
	newProduct := domain.Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}
	var updated bool = false
	var products = []domain.Product{}
	r.db.Read(&products)
	for i := range products {
		if products[i].Id == id {
			newProduct.Id = id
			products[i] = newProduct
			updated = true
		}
	}
	if !updated {
		return domain.Product{}, errorIdNotFound
	}
	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}
	return newProduct, nil
}

func (r *repository) Delete(id int) error {
	var pos int = -1
	var products = []domain.Product{}
	r.db.Read(&products)
	for i := range products {
		if products[i].Id == id {
			pos = i
		}
	}
	if pos == -1 {
		return errorIdNotFound
	}
	products = append(products[:pos], products[pos+1:]...)
	if err := r.db.Write(products); err != nil {
		return err
	}
	return nil
}

func (r *repository) UpdateMany(id int, name string, price int) (domain.Product, error) {
	var updated bool = false
	var product domain.Product
	var products = []domain.Product{}
	r.db.Read(&products)
	for i := range products {
		if products[i].Id == id {
			products[i].Name = name
			products[i].Price = price
			updated = true
			product = products[i]
		}
	}

	if !updated {
		return domain.Product{}, errorIdNotFound
	}
	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
