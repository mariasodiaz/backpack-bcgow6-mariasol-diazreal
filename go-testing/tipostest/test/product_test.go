package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/cmd/server/handler"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/internal/products"
	"github.com/mariasodiaz/backpack-bcgow6-mariasol-diazreal/go-testing/tipostest/pkg/store"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {

	_ = os.Setenv("TOKEN", "12345")
	store := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(store)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	r.PUT("/products/:id", p.Update())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "12345")

	return req, httptest.NewRecorder()
}

func TestUpdate(t *testing.T) {
	server := createServer()

	req, rr := createRequestTest(http.MethodPut, "/products/1", `{"id": 1, "name": "Tablet", "color": "Blanco", "price": 70000, "stock": 5, "code": "AL324", "published": true, "date": "20/05/20"}`)
	server.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
