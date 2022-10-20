package products

import (
	"errors"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/pkg/store"
)

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, name string, color string, price int, stock int, code string, published bool, date string) (Product, error)
	LastId() (int, error)
	Update(id int, name string, color string, price int, stock int, code string, published bool, date string) (Product, error)
	Delete(id int) error
	UpdateMany(id int, name string, price int) (Product, error)
}

type repository struct {
	db store.Store
}

type Product struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	Code      string `json:"code"`
	Published bool   `json:"published"`
	Date      string `json:"date"`
}

var errorIdNotFound = errors.New("id not found")

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var products = []Product{}
	r.db.Read(&products)
	return products, nil
}

func (r *repository) LastId() (int, error) {
	var products = []Product{}
	if err := r.db.Read(&products); err != nil {
		return 0, err
	}
	if len(products) == 0 {
		return 0, nil
	}
	return products[len(products)-1].Id, nil
}

func (r *repository) Store(id int, name string, color string, price int, stock int, code string, published bool, date string) (Product, error) {
	product := Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}
	var products = []Product{}
	r.db.Read(&products)
	products = append(products, product)
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return product, nil
}

func (r *repository) Update(id int, name string, color string, price int, stock int, code string, published bool, date string) (Product, error) {
	newProduct := Product{Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}
	var updated bool = false
	var products = []Product{}
	r.db.Read(&products)
	for i := range products {
		if products[i].Id == id {
			newProduct.Id = id
			products[i] = newProduct
			updated = true
		}
	}
	if !updated {
		return Product{}, errorIdNotFound
	}
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return newProduct, nil
}

func (r *repository) Delete(id int) error {
	var pos int = -1
	var products = []Product{}
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

func (r *repository) UpdateMany(id int, name string, price int) (Product, error) {
	var updated bool = false
	var product Product
	var products = []Product{}
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
		return Product{}, errorIdNotFound
	}
	if err := r.db.Write(products); err != nil {
		return Product{}, err
	}
	return product, nil
}
