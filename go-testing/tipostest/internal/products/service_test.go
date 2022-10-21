package products

import (
	"testing"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/internal/domain"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestMockUpdateMany(t *testing.T) {
	products := []domain.Product{{Id: 1, Name: "Computadora", Price: 150000}}
	mock := store.MockStorage{Products: products}
	repository := NewRepository(&mock)
	service := NewService(repository)
	newProduct := &domain.Product{
		Id:    1,
		Name:  "Televisor",
		Price: 12,
	}
	product, err := service.UpdateMany(1, "Televisor", 12)
	assert.Nil(t, err)
	assert.Equal(t, newProduct.Name, product.Name)
}

func TestMockDelete(t *testing.T) {
	products := []domain.Product{{Id: 1, Name: "Computadora", Price: 150000}, {Id: 2, Name: "Televisor", Price: 10000}}
	mock := store.MockStorage{Products: products}
	repository := NewRepository(&mock)
	service := NewService(repository)
	err := service.Delete(1)

	assert.Nil(t, err)
	err = service.Delete(28)
	assert.NotNil(t, err)

}

func TestMockGetAll(t *testing.T) {
	products := []domain.Product{{Id: 1, Name: "Computadora", Price: 150000}, {Id: 2, Name: "Televisor", Price: 10000}}
	mock := store.MockStorage{Products: products}
	repository := NewRepository(&mock)
	service := NewService(repository)
	result, err := service.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, products, result)
}

func TestMockStore(t *testing.T) {
	product := domain.Product{Id: 1, Name: "Computadora", Color: "Negro", Price: 150000, Stock: 5, Code: "AK324", Published: true, Date: "20/05/20"}
	mock := store.MockStorage{}
	repository := NewRepository(&mock)
	service := NewService(repository)
	result, err := service.Store(product.Name, product.Color, product.Price, product.Stock, product.Code, product.Published, product.Date)
	assert.Nil(t, err)
	assert.Equal(t, product, result)
}

func TestMockUpdate(t *testing.T) {
	product := []domain.Product{{Id: 1, Name: "Computadora", Color: "Negro", Price: 150000, Stock: 5, Code: "AK324", Published: true, Date: "20/05/20"}}
	productUpdated := domain.Product{Id: 1, Name: "Televisor", Color: "Gris", Price: 2000000, Stock: 3, Code: "ABB324", Published: false, Date: "22/05/20"}
	mock := store.MockStorage{Products: product}
	repository := NewRepository(&mock)
	service := NewService(repository)
	result, err := service.Update(productUpdated.Id, productUpdated.Name, productUpdated.Color, productUpdated.Price, productUpdated.Stock, productUpdated.Code, productUpdated.Published, productUpdated.Date)
	assert.Nil(t, err)
	assert.Equal(t, productUpdated, result)
}
