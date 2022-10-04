package products

import "errors"

type Repository interface {
	GetAll() ([]Product, error)
	GetById(id int) (Product, error)
	Store(id int, name string, color string, price int, stock int, code string, published bool, date string) (Product, error)
	LastId() (int, error)
}

type repository struct {
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

var products = []Product{
	{Id: 1, Name: "Cama", Color: "Blanco", Price: 150000, Stock: 5, Code: "AF289A", Published: true, Date: "20/09/2022"},
	{Id: 2, Name: "Televisor", Color: "Negro", Price: 200000, Stock: 1, Code: "AF289A", Published: false, Date: "22/09/2022"},
	{Id: 3, Name: "Cocina", Color: "Plateado", Price: 80000, Stock: 3, Code: "BE224A", Published: true, Date: "01/01/2022"},
	{Id: 4, Name: "Plancha", Color: "Blanco", Price: 200000, Stock: 10, Code: "CL208D", Published: true, Date: "15/04/2022"},
}
var lastId int = 4

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return products, nil
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}

func (r *repository) Store(id int, name string, color string, price int, stock int, code string, published bool, date string) (Product, error) {
	lastId = id
	product := Product{Id: id, Name: name, Color: color, Price: price, Stock: stock, Code: code, Published: published, Date: date}
	products = append(products, product)
	return product, nil
}

func (r *repository) GetById(id int) (Product, error) {
	for _, value := range products {
		if value.Id == id {
			return value, nil
		}
	}
	return Product{}, errors.New("id no encontrado")
}
