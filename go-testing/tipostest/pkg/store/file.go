package store

import (
	"encoding/json"
	"os"

	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/internal/domain"
)

type Type string

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type MockStorage struct {
	ReadWasCalled bool
	Products      []domain.Product
	Error         error
}

const (
	FileType Type = "file"
)

type fileStore struct {
	FilePath string
}

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &fileStore{fileName}
	}
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	m.ReadWasCalled = true
	a := data.(*[]domain.Product)
	*a = m.Products
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	a := data.([]domain.Product)
	m.Products = append(m.Products, a[len(a)-1])
	return nil
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
