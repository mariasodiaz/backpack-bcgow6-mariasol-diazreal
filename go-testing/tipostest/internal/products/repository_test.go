package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct {
	Products []Product
}

type MockDB struct {
	ReadWasCalled bool
	BeforeUpdate  Product
	AfterUpdate   Product
	Products      []Product
}

func (s StubDB) Read(data interface{}) error {
	a := data.(*[]Product)
	*a = s.Products
	return nil
}

func (s StubDB) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	products := []Product{
		{
			Id:    1,
			Name:  "Computadora",
			Color: "Gris",
			Price: 120000,
			Stock: 2,
			Code:  "AF662",
		}, {
			Id:    2,
			Name:  "Televisor",
			Color: "Negro",
			Price: 300000,
			Stock: 1,
			Code:  "AB769",
		},
	}

	myStubDB := StubDB{products}
	motor := NewRepository(&myStubDB)

	result, err := motor.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, products, result)

}

func (m *MockDB) Read(data interface{}) error {
	m.ReadWasCalled = true
	a := data.(*[]Product)
	*a = m.Products
	return nil
}

func (m *MockDB) Write(data interface{}) error {
	a := data.([]Product)
	m.Products = append(m.Products, a[len(a)-1])
	return nil
}

func TestUpdate(t *testing.T) {
	Update := Product{Id: 1, Name: "Computadora", Price: 230000}
	Updated := Product{Id: 1, Name: "Televisor", Price: 180000}
	products := []Product{Update}
	myMockDB := MockDB{ReadWasCalled: false, BeforeUpdate: Update, AfterUpdate: Updated, Products: products}
	motor := NewRepository(&myMockDB)

	productUpdated, err := motor.UpdateMany(Update.Id, Updated.Name, Updated.Price)

	assert.Nil(t, err)
	assert.Equal(t, Updated, productUpdated)
	assert.True(t, myMockDB.ReadWasCalled)

}
